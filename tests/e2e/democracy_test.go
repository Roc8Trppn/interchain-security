package e2e_test

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	proposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	appConsumer "github.com/cosmos/interchain-security/app/consumer-democracy"
	appProvider "github.com/cosmos/interchain-security/app/provider"
	"github.com/cosmos/interchain-security/testutil/simapp"
	consumerkeeper "github.com/cosmos/interchain-security/x/ccv/consumer/keeper"
	consumertypes "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	"github.com/cosmos/interchain-security/x/ccv/types"
	"github.com/cosmos/interchain-security/x/ccv/utils"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/stretchr/testify/suite"
)

var consumerFraction, _ = sdk.NewDecFromStr(consumerkeeper.ConsumerRedistributeFrac)

type ConsumerDemocracyTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains
	providerChain *ibctesting.TestChain
	consumerChain *ibctesting.TestChain

	path         *ibctesting.Path
	transferPath *ibctesting.Path
}

func (s *ConsumerDemocracyTestSuite) SetupTest() {
	s.coordinator, s.providerChain, s.consumerChain = simapp.NewProviderConsumerDemocracyCoordinator(s.T())

	// valsets must match
	providerValUpdates := tmtypes.TM2PB.ValidatorUpdates(s.providerChain.Vals)
	consumerValUpdates := tmtypes.TM2PB.ValidatorUpdates(s.consumerChain.Vals)
	s.Require().True(len(providerValUpdates) == len(consumerValUpdates), "initial valset not matching")
	for i := 0; i < len(providerValUpdates); i++ {
		addr1 := utils.GetChangePubKeyAddress(providerValUpdates[i])
		addr2 := utils.GetChangePubKeyAddress(consumerValUpdates[i])
		s.Require().True(bytes.Equal(addr1, addr2), "validator mismatch")
	}

	// move both chains to the next block
	s.providerChain.NextBlock()
	s.consumerChain.NextBlock()

	// create consumer client on provider chain and set as consumer client for consumer chainID in provider keeper.
	err := s.providerChain.App.(*appProvider.App).ProviderKeeper.CreateConsumerClient(
		s.providerCtx(),
		s.consumerChain.ChainID,
		s.consumerChain.LastHeader.GetHeight().(clienttypes.Height),
		false,
	)
	s.Require().NoError(err)

	// move provider to next block to commit the state
	s.providerChain.NextBlock()

	// initialize the consumer chain with the genesis state stored on the provider
	consumerGenesis, found := s.providerChain.App.(*appProvider.App).ProviderKeeper.GetConsumerGenesis(
		s.providerCtx(),
		s.consumerChain.ChainID,
	)
	s.Require().True(found, "consumer genesis not found")
	s.consumerChain.App.(*appConsumer.App).ConsumerKeeper.InitGenesis(s.consumerChain.GetContext(), &consumerGenesis)

	// create path for the CCV channel
	s.path = ibctesting.NewPath(s.consumerChain, s.providerChain)

	// update CCV path with correct info
	// - set provider endpoint's clientID
	consumerClient, found := s.providerChain.App.(*appProvider.App).ProviderKeeper.GetConsumerClientId(
		s.providerCtx(),
		s.consumerChain.ChainID,
	)
	s.Require().True(found, "consumer client not found")
	s.path.EndpointB.ClientID = consumerClient
	// - set consumer endpoint's clientID
	providerClient, found := s.consumerChain.App.(*appConsumer.App).ConsumerKeeper.GetProviderClientID(s.consumerChain.GetContext())
	s.Require().True(found, "provider client not found")
	s.path.EndpointA.ClientID = providerClient
	// - client config
	providerUnbondingPeriod := s.providerChain.App.(*appProvider.App).GetStakingKeeper().UnbondingTime(s.providerCtx())
	s.path.EndpointB.ClientConfig.(*ibctesting.TendermintConfig).UnbondingPeriod = providerUnbondingPeriod
	s.path.EndpointB.ClientConfig.(*ibctesting.TendermintConfig).TrustingPeriod = providerUnbondingPeriod / utils.TrustingPeriodFraction
	consumerUnbondingPeriod := utils.ComputeConsumerUnbondingPeriod(providerUnbondingPeriod)
	s.path.EndpointA.ClientConfig.(*ibctesting.TendermintConfig).UnbondingPeriod = consumerUnbondingPeriod
	s.path.EndpointA.ClientConfig.(*ibctesting.TendermintConfig).TrustingPeriod = consumerUnbondingPeriod / utils.TrustingPeriodFraction
	// - channel config
	s.path.EndpointA.ChannelConfig.PortID = types.ConsumerPortID
	s.path.EndpointB.ChannelConfig.PortID = types.ProviderPortID
	s.path.EndpointA.ChannelConfig.Version = types.Version
	s.path.EndpointB.ChannelConfig.Version = types.Version
	s.path.EndpointA.ChannelConfig.Order = channeltypes.ORDERED
	s.path.EndpointB.ChannelConfig.Order = channeltypes.ORDERED

	// set chains sender account number
	// TODO: to be fixed in #151
	err = s.path.EndpointB.Chain.SenderAccount.SetAccountNumber(6)
	s.Require().NoError(err)
	err = s.path.EndpointA.Chain.SenderAccount.SetAccountNumber(0)
	s.Require().NoError(err)

	// create path for the transfer channel
	s.transferPath = ibctesting.NewPath(s.consumerChain, s.providerChain)
	s.transferPath.EndpointA.ChannelConfig.PortID = transfertypes.PortID
	s.transferPath.EndpointB.ChannelConfig.PortID = transfertypes.PortID
	s.transferPath.EndpointA.ChannelConfig.Version = transfertypes.Version
	s.transferPath.EndpointB.ChannelConfig.Version = transfertypes.Version
}

func (s *ConsumerDemocracyTestSuite) SetupCCVChannel() {
	s.StartSetupCCVChannel()
	s.CompleteSetupCCVChannel()
	s.SetupTransferChannel()
}

func (s *ConsumerDemocracyTestSuite) StartSetupCCVChannel() {
	s.coordinator.CreateConnections(s.path)

	err := s.path.EndpointA.ChanOpenInit()
	s.Require().NoError(err)

	err = s.path.EndpointB.ChanOpenTry()
	s.Require().NoError(err)
}

func (s *ConsumerDemocracyTestSuite) CompleteSetupCCVChannel() {
	err := s.path.EndpointA.ChanOpenAck()
	s.Require().NoError(err)

	err = s.path.EndpointB.ChanOpenConfirm()
	s.Require().NoError(err)

	// ensure counterparty is up to date
	err = s.path.EndpointA.UpdateClient()
	s.Require().NoError(err)
}

func (s *ConsumerDemocracyTestSuite) SetupTransferChannel() {
	// transfer path will use the same connection as ccv path

	s.transferPath.EndpointA.ClientID = s.path.EndpointA.ClientID
	s.transferPath.EndpointA.ConnectionID = s.path.EndpointA.ConnectionID
	s.transferPath.EndpointB.ClientID = s.path.EndpointB.ClientID
	s.transferPath.EndpointB.ConnectionID = s.path.EndpointB.ConnectionID

	// CCV channel handshake will automatically initiate transfer channel handshake on ACK
	// so transfer channel will be on stage INIT when CompleteSetupCCVChannel returns.
	s.transferPath.EndpointA.ChannelID = s.consumerChain.App.(*appConsumer.App).
		ConsumerKeeper.GetDistributionTransmissionChannel(s.consumerChain.GetContext())

	// Complete TRY, ACK, CONFIRM for transfer path
	err := s.transferPath.EndpointB.ChanOpenTry()
	s.Require().NoError(err)

	err = s.transferPath.EndpointA.ChanOpenAck()
	s.Require().NoError(err)

	err = s.transferPath.EndpointB.ChanOpenConfirm()
	s.Require().NoError(err)

	// ensure counterparty is up to date
	err = s.transferPath.EndpointA.UpdateClient()
	s.Require().NoError(err)
}

func TestConsumerDemocracyTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerDemocracyTestSuite))
}

func (s *ConsumerDemocracyTestSuite) TestDemocracyRewarsDistribution() {

	s.consumerChain.NextBlock()
	stakingKeeper := s.consumerChain.App.(*appConsumer.App).StakingKeeper
	authKeeper := s.consumerChain.App.(*appConsumer.App).AccountKeeper
	distrKeeper := s.consumerChain.App.(*appConsumer.App).DistrKeeper
	bankKeeper := s.consumerChain.App.(*appConsumer.App).BankKeeper
	bondDenom := stakingKeeper.BondDenom(s.consumerCtx())

	currentRepresentativesRewards := map[string]sdk.Dec{}
	nextRepresentativesRewards := map[string]sdk.Dec{}
	representativesTokens := map[string]sdk.Int{}

	for _, representative := range stakingKeeper.GetAllValidators(s.consumerCtx()) {
		currentRepresentativesRewards[representative.OperatorAddress] = sdk.NewDec(0)
		nextRepresentativesRewards[representative.OperatorAddress] = sdk.NewDec(0)
		representativesTokens[representative.OperatorAddress] = representative.GetTokens()
	}

	distrModuleAccount := distrKeeper.GetDistributionAccount(s.consumerCtx())
	providerRedistributeAccount := authKeeper.GetModuleAccount(s.consumerCtx(), consumertypes.ConsumerToSendToProviderName)
	//balance of consumer redistribute address will always be 0 when checked between 2 NextBlock() calls

	currentDistrModuleAccountBalance := sdk.NewDecFromInt(bankKeeper.GetBalance(s.consumerCtx(), distrModuleAccount.GetAddress(), bondDenom).Amount)
	currentProviderFeeAccountBalance := sdk.NewDecFromInt(bankKeeper.GetBalance(s.consumerCtx(), providerRedistributeAccount.GetAddress(), bondDenom).Amount)
	currentCommunityPoolBalance := distrKeeper.GetFeePoolCommunityCoins(s.consumerCtx()).AmountOf(bondDenom)
	for key := range currentRepresentativesRewards {
		representativeAddr, _ := sdk.ValAddressFromBech32(key)
		representativeReward := distrKeeper.GetValidatorOutstandingRewards(s.consumerCtx(), representativeAddr).Rewards.AmountOf(bondDenom)
		currentRepresentativesRewards[key] = representativeReward
	}

	s.consumerChain.NextBlock()

	nextDistrModuleAccountBalance := sdk.NewDecFromInt(bankKeeper.GetBalance(s.consumerCtx(), distrModuleAccount.GetAddress(), bondDenom).Amount)
	nextProviderFeeAccountBalance := sdk.NewDecFromInt(bankKeeper.GetBalance(s.consumerCtx(), providerRedistributeAccount.GetAddress(), bondDenom).Amount)
	nextCommunityPoolBalance := distrKeeper.GetFeePoolCommunityCoins(s.consumerCtx()).AmountOf(bondDenom)
	for key := range nextRepresentativesRewards {
		representativeAddr, _ := sdk.ValAddressFromBech32(key)
		representativeReward := distrKeeper.GetValidatorOutstandingRewards(s.consumerCtx(), representativeAddr).Rewards.AmountOf(bondDenom)
		nextRepresentativesRewards[key] = representativeReward
	}

	distrModuleDifference := nextDistrModuleAccountBalance.Sub(currentDistrModuleAccountBalance)
	providerDifference := nextProviderFeeAccountBalance.Sub(currentProviderFeeAccountBalance)
	communityPoolDifference := nextCommunityPoolBalance.Sub(currentCommunityPoolBalance)
	representativeDifference := map[string]sdk.Dec{}
	consumerRedistributeDifference := communityPoolDifference

	for key, currentReward := range currentRepresentativesRewards {
		representativeDifference[key] = nextRepresentativesRewards[key].Sub(currentReward)
		consumerRedistributeDifference = consumerRedistributeDifference.Add(representativeDifference[key])
	}

	//confirm that the total amount given to the community pool plus all representatives is equal to the total amount taken out of distribution
	s.Require().Equal(distrModuleDifference, consumerRedistributeDifference)
	//confirm that the percentage given to the community pool is equal to the configured community tax percentage.
	s.Require().Equal(communityPoolDifference.Quo(consumerRedistributeDifference), distrKeeper.GetCommunityTax(s.consumerCtx()))
	//check that the fraction actually kept by the consumer is the correct fraction. using InEpsilon because the math code uses truncations
	s.Require().InEpsilon(distrModuleDifference.Quo(providerDifference.Add(distrModuleDifference)).MustFloat64(), consumerFraction.MustFloat64(), float64(0.0001))
	//check that the fraction actually kept by the provider is the correct fraction. using InEpsilon because the math code uses truncations
	s.Require().InEpsilon(providerDifference.Quo(providerDifference.Add(distrModuleDifference)).MustFloat64(), sdk.NewDec(1).Sub(consumerFraction).MustFloat64(), float64(0.0001))

	totalRepresentativePower := stakingKeeper.GetValidatorSet().TotalBondedTokens(s.consumerCtx())

	//check that each representative has gotten the correct amount of rewards
	for key, representativeTokens := range representativesTokens {
		powerFraction := sdk.NewDecFromInt(representativeTokens).QuoTruncate(sdk.NewDecFromInt(totalRepresentativePower))
		s.Require().Equal(powerFraction, representativeDifference[key].Quo(consumerRedistributeDifference.Sub(communityPoolDifference)))
	}
}

func (s *ConsumerDemocracyTestSuite) TestDemocracyGovernanceWhitelisting() {
	govKeeper := s.consumerChain.App.(*appConsumer.App).GovKeeper
	stakingKeeper := s.consumerChain.App.(*appConsumer.App).StakingKeeper
	bankKeeper := s.consumerChain.App.(*appConsumer.App).BankKeeper
	authKeeper := s.consumerChain.App.(*appConsumer.App).AccountKeeper
	mintKeeper := s.consumerChain.App.(*appConsumer.App).MintKeeper
	newAuthParamValue := uint64(128)
	newMintParamValue := sdk.NewDecWithPrec(1, 1) // "0.100000000000000000"
	allowedChange := proposal.ParamChange{Subspace: minttypes.ModuleName, Key: "InflationMax", Value: fmt.Sprintf("\"%s\"", newMintParamValue)}
	forbiddenChange := proposal.ParamChange{Subspace: authtypes.ModuleName, Key: "MaxMemoCharacters", Value: fmt.Sprintf("\"%s\"", strconv.FormatUint(newAuthParamValue, 10))}
	votingAccounts := s.consumerChain.SenderAccounts
	bondDenom := stakingKeeper.BondDenom(s.consumerCtx())
	depositAmount := govKeeper.GetDepositParams(s.consumerCtx()).MinDeposit
	votingParams := govKeeper.GetVotingParams(s.consumerCtx())
	votingParams.VotingPeriod = 3 * time.Second
	govKeeper.SetVotingParams(s.consumerCtx(), votingParams)
	s.consumerChain.NextBlock()
	votersOldBalances := getAccountsBalances(s.consumerCtx(), bankKeeper, bondDenom, votingAccounts)

	//submit proposal with forbidden and allowed changes
	paramChange := proposaltypes.ParameterChangeProposal{Changes: []proposaltypes.ParamChange{allowedChange, forbiddenChange}}
	err := submitProposalWithDepositAndVote(govKeeper, s.consumerCtx(), paramChange, votingAccounts, depositAmount)
	s.Assert().NoError(err)
	//set current header time to be equal or later than voting end time in order to process proposal from active queue,
	//once the proposal is added to the chain
	s.consumerChain.CurrentHeader.Time = s.consumerChain.CurrentHeader.Time.Add(votingParams.VotingPeriod)
	s.consumerChain.NextBlock()
	//at this moment, proposal is added, but not yet executed. we are saving old param values for comparison
	oldAuthParamValue := authKeeper.GetParams(s.consumerCtx()).MaxMemoCharacters
	oldMintParamValue := mintKeeper.GetParams(s.consumerCtx()).InflationMax
	s.consumerChain.NextBlock()
	//at this moment, proposal is executed or deleted if forbidden
	currentAuthParamValue := authKeeper.GetParams(s.consumerCtx()).MaxMemoCharacters
	currentMintParamValue := mintKeeper.GetParams(s.consumerCtx()).InflationMax
	//check that parameters are not changed, since the proposal contained both forbidden and allowed changes
	s.Assert().Equal(oldAuthParamValue, currentAuthParamValue)
	s.Assert().NotEqual(newAuthParamValue, currentAuthParamValue)
	s.Assert().Equal(oldMintParamValue, currentMintParamValue)
	s.Assert().NotEqual(newMintParamValue, currentMintParamValue)
	//deposit is refunded
	s.Assert().Equal(votersOldBalances, getAccountsBalances(s.consumerCtx(), bankKeeper, bondDenom, votingAccounts))

	//submit proposal with allowed changes
	paramChange = proposaltypes.ParameterChangeProposal{Changes: []proposaltypes.ParamChange{allowedChange}}
	err = submitProposalWithDepositAndVote(govKeeper, s.consumerCtx(), paramChange, votingAccounts, depositAmount)
	s.Assert().NoError(err)
	s.consumerChain.CurrentHeader.Time = s.consumerChain.CurrentHeader.Time.Add(votingParams.VotingPeriod)
	s.consumerChain.NextBlock()
	oldMintParamValue = mintKeeper.GetParams(s.consumerCtx()).InflationMax
	s.consumerChain.NextBlock()
	currentMintParamValue = mintKeeper.GetParams(s.consumerCtx()).InflationMax
	//check that parameters are changed, since the proposal contained only allowed changes
	s.Assert().Equal(newMintParamValue, currentMintParamValue)
	s.Assert().NotEqual(oldMintParamValue, currentMintParamValue)
	//deposit is refunded
	s.Assert().Equal(votersOldBalances, getAccountsBalances(s.consumerCtx(), bankKeeper, bondDenom, votingAccounts))

	//submit proposal with forbidden changes
	paramChange = proposaltypes.ParameterChangeProposal{Changes: []proposaltypes.ParamChange{forbiddenChange}}
	err = submitProposalWithDepositAndVote(govKeeper, s.consumerCtx(), paramChange, votingAccounts, depositAmount)
	s.Assert().NoError(err)
	s.consumerChain.CurrentHeader.Time = s.consumerChain.CurrentHeader.Time.Add(votingParams.VotingPeriod)
	s.consumerChain.NextBlock()
	oldAuthParamValue = authKeeper.GetParams(s.consumerCtx()).MaxMemoCharacters
	s.consumerChain.NextBlock()
	currentAuthParamValue = authKeeper.GetParams(s.consumerCtx()).MaxMemoCharacters
	//check that parameters are not changed, since the proposal contained forbidden changes
	s.Assert().Equal(oldAuthParamValue, currentAuthParamValue)
	s.Assert().NotEqual(newAuthParamValue, currentAuthParamValue)
	//deposit is refunded
	s.Assert().Equal(votersOldBalances, getAccountsBalances(s.consumerCtx(), bankKeeper, bondDenom, votingAccounts))
}

func submitProposalWithDepositAndVote(govKeeper govkeeper.Keeper, ctx sdk.Context, paramChange proposaltypes.ParameterChangeProposal,
	accounts []ibctesting.SenderAccount, depositAmount sdk.Coins) error {
	proposal, err := govKeeper.SubmitProposal(ctx, &paramChange)
	if err != nil {
		return err
	}
	_, err = govKeeper.AddDeposit(ctx, proposal.ProposalId, accounts[0].SenderAccount.GetAddress(), depositAmount) //proposal becomes active
	if err != nil {
		return err
	}

	for _, account := range accounts {
		err = govKeeper.AddVote(ctx, proposal.ProposalId, account.SenderAccount.GetAddress(), govtypes.NewNonSplitVoteOption(govtypes.OptionYes))
		if err != nil {
			return err
		}
	}
	return nil
}

func getAccountsBalances(ctx sdk.Context, bankKeeper bankkeeper.Keeper, bondDenom string, accounts []ibctesting.SenderAccount) map[string]sdk.Int {
	accountsBalances := map[string]sdk.Int{}
	for _, acc := range accounts {
		accountsBalances[string(acc.SenderAccount.GetAddress())] =
			bankKeeper.GetBalance(ctx, acc.SenderAccount.GetAddress(), bondDenom).Amount
	}

	return accountsBalances
}

func (s *ConsumerDemocracyTestSuite) providerCtx() sdk.Context {
	return s.providerChain.GetContext()
}

func (s *ConsumerDemocracyTestSuite) consumerCtx() sdk.Context {
	return s.consumerChain.GetContext()
}