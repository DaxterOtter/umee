package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"

	gravity "github.com/Gravity-Bridge/Gravity-Bridge/module/x/gravity"
	gravitykeeper "github.com/Gravity-Bridge/Gravity-Bridge/module/x/gravity/keeper"
	gravitytypes "github.com/Gravity-Bridge/Gravity-Bridge/module/x/gravity/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfer "github.com/cosmos/ibc-go/v2/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v2/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v2/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v2/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v2/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v2/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/osmosis-labs/bech32-ibc/x/bech32ibc"
	bech32ibckeeper "github.com/osmosis-labs/bech32-ibc/x/bech32ibc/keeper"
	bech32ibctypes "github.com/osmosis-labs/bech32-ibc/x/bech32ibc/types"
	customante "github.com/umee-network/umee/v2/ante"
	appparams "github.com/umee-network/umee/v2/app/params"
	"github.com/umee-network/umee/v2/app/upgrades"
	"github.com/umee-network/umee/v2/app/upgrades/calypso"
	"github.com/umee-network/umee/v2/app/upgrades/cosmwasm"
	uibctransfer "github.com/umee-network/umee/v2/x/ibctransfer"
	uibctransferkeeper "github.com/umee-network/umee/v2/x/ibctransfer/keeper"
	"github.com/umee-network/umee/v2/x/leverage"
	leverageclient "github.com/umee-network/umee/v2/x/leverage/client"
	leveragekeeper "github.com/umee-network/umee/v2/x/leverage/keeper"
	leveragetypes "github.com/umee-network/umee/v2/x/leverage/types"
	"github.com/umee-network/umee/v2/x/oracle"
	oraclekeeper "github.com/umee-network/umee/v2/x/oracle/keeper"
	oracletypes "github.com/umee-network/umee/v2/x/oracle/types"
)

const (
	// Name defines the application name of the Umee network.
	Name = "umee"

	// BondDenom defines the native staking token denomination.
	BondDenom = "uumee"

	// DisplayDenom defines the name, symbol, and display value of the umee token.
	DisplayDenom = "UMEE"

	// AddrLen is the size allowed length (in bytes) for an address.
	//
	// NOTE: In the SDK, the default value is 255.
	AddrLen = 20

	// ContractAddrLen is the size allowed length (in bytes) for an contract address.
	ContractAddrLen = 32
)

var (
	_ CosmosApp               = (*UmeeApp)(nil)
	_ servertypes.Application = (*UmeeApp)(nil)

	// DefaultNodeHome defines the default home directory for the application
	// daemon.
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		GenutilModule{},
		BankModule{},
		capability.AppModuleBasic{},
		StakingModule{},
		MintModule{},
		distr.AppModuleBasic{},
		GovModule{AppModuleBasic: gov.NewAppModuleBasic(getGovProposalHandlers()...)},
		params.AppModuleBasic{},
		CrisisModule{},
		SlashingModule{},
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		gravity.AppModuleBasic{},
		leverage.AppModuleBasic{},
		oracle.AppModuleBasic{},
		bech32ibc.AppModuleBasic{},
		wasm.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		gravitytypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
		leveragetypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		oracletypes.ModuleName:         nil,
		wasm.ModuleName:                {authtypes.Burner},
	}
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("failed to get user home directory: %s", err))
	}

	DefaultNodeHome = filepath.Join(userHomeDir, fmt.Sprintf(".%s", Name))

	// XXX: If other upstream or external application's depend on any of Umee's
	// CLI or command functionality, then this would require us to move the
	// SetAddressConfig call to somewhere external such as the root command
	// constructor and anywhere else we contract the app.
	SetAddressConfig()
}

// UmeeApp defines the ABCI application for the Umee network as an extension of
// the Cosmos SDK's BaseApp.
type UmeeApp struct {
	*baseapp.BaseApp

	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   uibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper
	GravityKeeper    gravitykeeper.Keeper
	LeverageKeeper   leveragekeeper.Keeper
	OracleKeeper     oraclekeeper.Keeper
	bech32IbcKeeper  bech32ibckeeper.Keeper

	// make scoped keepers public for testing purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// configurator
	configurator *module.Configurator

	// CosmWasm
	WasmKeeper       wasm.Keeper
	scopedWasmKeeper capabilitykeeper.ScopedKeeper
}

func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig appparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *UmeeApp {
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	base := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	base.SetCommitMultiStoreTracer(traceStore)
	base.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		feegrant.StoreKey, authzkeeper.StoreKey, gravitytypes.StoreKey,
		leveragetypes.StoreKey, oracletypes.StoreKey, bech32ibctypes.StoreKey,
		wasm.StoreKey,
	)
	transientKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &UmeeApp{
		BaseApp:           base,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             transientKeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		transientKeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	base.SetParamStore(
		app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	app.ScopedIBCKeeper = app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	app.ScopedTransferKeeper = app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.scopedWasmKeeper = app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)

	// Applications that wish to enforce statically created ScopedKeepers should
	// call `Seal` after creating their scoped modules in the app via
	// `ScopeToModule`.
	app.CapabilityKeeper.Seal()

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)
	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.BaseApp.MsgServiceRouter(),
	)
	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.ModuleAccountAddrs(),
	)
	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&app.StakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&app.StakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
	)
	app.OracleKeeper = oraclekeeper.NewKeeper(
		appCodec,
		keys[oracletypes.ModuleName],
		app.GetSubspace(oracletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.DistrKeeper,
		&app.StakingKeeper,
		distrtypes.ModuleName,
	)
	app.LeverageKeeper = leveragekeeper.NewKeeper(
		appCodec,
		keys[leveragetypes.ModuleName],
		app.GetSubspace(leveragetypes.ModuleName),
		app.BankKeeper,
		app.OracleKeeper,
	)
	app.LeverageKeeper = *app.LeverageKeeper.SetHooks(
		leveragetypes.NewMultiHooks(
			app.OracleKeeper.Hooks(),
		),
	)

	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		app.ScopedIBCKeeper,
	)

	// Create an original ICS-20 transfer keeper and AppModule and then use it to
	// created an Umee wrapped ICS-20 transfer keeper and AppModule.
	ibcTransferKeeper := ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.ScopedTransferKeeper,
	)
	app.TransferKeeper = uibctransferkeeper.New(ibcTransferKeeper, app.BankKeeper)
	transferModule := uibctransfer.NewAppModule(ibctransfer.NewAppModule(ibcTransferKeeper), app.TransferKeeper)

	baseBankKeeper := app.BankKeeper.(bankkeeper.BaseKeeper)

	app.bech32IbcKeeper = *bech32ibckeeper.NewKeeper(
		app.IBCKeeper.ChannelKeeper, appCodec, keys[bech32ibctypes.StoreKey],
		app.TransferKeeper,
	)

	app.GravityKeeper = gravitykeeper.NewKeeper(
		keys[gravitytypes.StoreKey],
		app.GetSubspace(gravitytypes.ModuleName),
		appCodec,
		&baseBankKeeper,
		&app.StakingKeeper,
		&app.SlashingKeeper,
		&app.DistrKeeper,
		&app.AccountKeeper,
		&app.TransferKeeper.Keeper,
		&app.bech32IbcKeeper,
	)

	wasmDir := filepath.Join(homePath, "data")

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := "iterator,staking,stargate"
	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		GetWasmOpts(appOpts)...,
	)

	// register the staking hooks
	//
	// NOTE: The stakingKeeper above is passed by reference, so that it will contain
	// these hooks.
	app.StakingKeeper = *app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
			app.GravityKeeper.Hooks(),
		),
	)

	// create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
	ibcRouter.AddRoute(wasm.ModuleName, wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper))
	app.IBCKeeper.SetRouter(ibcRouter)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(leveragetypes.RouterKey, leverage.NewUpdateRegistryProposalHandler(app.LeverageKeeper)).
		AddRoute(gravitytypes.RouterKey, gravitykeeper.NewGravityProposalHandler(app.GravityKeeper)).
		AddRoute(bech32ibctypes.RouterKey, bech32ibc.NewBech32IBCProposalHandler(app.bech32IbcKeeper))

	// register wasm gov proposal types
	enabledProposals := GetWasmEnabledProposals()
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	}

	// Create evidence Keeper so we can register the IBC light client misbehavior
	// evidence route.
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)

	app.EvidenceKeeper = *evidenceKeeper
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&app.StakingKeeper,
		govRouter,
	)

	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	app.registerStoreLoaders()

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		gravity.NewAppModule(app.GravityKeeper, app.BankKeeper),
		leverage.NewAppModule(appCodec, app.LeverageKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		bech32ibc.NewAppModule(appCodec, app.bech32IbcKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that there
	// is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	//
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		ibctransfertypes.ModuleName,
		vestingtypes.ModuleName,
		banktypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		govtypes.ModuleName,
		authtypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
		leveragetypes.ModuleName,
		oracletypes.ModuleName,
		gravitytypes.ModuleName,
		bech32ibctypes.ModuleName,
		wasm.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		banktypes.ModuleName,
		capabilitytypes.ModuleName,
		vestingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		upgradetypes.ModuleName,
		ibchost.ModuleName,
		distrtypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		authtypes.ModuleName,
		minttypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		leveragetypes.ModuleName,
		oracletypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		stakingtypes.ModuleName,
		gravitytypes.ModuleName,
		bech32ibctypes.ModuleName,
		wasm.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	//
	// NOTE: Capability module must occur first so that it can initialize any
	// capabilities so that other modules that want to create or claim capabilities
	// afterwards in InitChain can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		upgradetypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		oracletypes.ModuleName,
		leveragetypes.ModuleName,
		gravitytypes.ModuleName,
		bech32ibctypes.ModuleName,
		wasm.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	configurator := module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.configurator = &configurator
	app.mm.RegisterServices(*app.configurator)

	// Create the simulation manager and define the order of the modules for
	// deterministic simulations.
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		// TODO: Ensure x/leverage implements simulator and then uncomment.
		// leverage.NewAppModule(appCodec, app.LeverageKeeper, app.AccountKeeper, app.BankKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(transientKeys)
	app.MountMemoryStores(memKeys)

	anteHandler, err := customante.NewAnteHandler(
		customante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			FeegrantKeeper:  app.FeeGrantKeeper,
			OracleKeeper:    app.OracleKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			WasmConfig:      wasmConfig,
			WasmTxCounter:   keys[wasm.StoreKey],
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create ante handler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	upgrades.RegisterUpgradeHandlers(
		app.mm, app.configurator, &app.AccountKeeper, &baseBankKeeper, &app.bech32IbcKeeper, &app.DistrKeeper,
		&app.MintKeeper, &app.StakingKeeper, &app.UpgradeKeeper, &app.LeverageKeeper, &app.OracleKeeper,
		&app.WasmKeeper,
	)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(fmt.Sprintf("failed to load latest version: %s", err))
		}
	}

	return app
}

// Sets up the StoreLoader for new, deleted, or renamed modules
func (app *UmeeApp) registerStoreLoaders() {
	// Read the upgrade height and name from previous execution
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	// v1->Calypso STORE LOADER SETUP
	// Register the new Calypso modules and the special StoreLoader to add them
	if upgradeInfo.Name == calypso.PlanName {
		if !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) { // Recognized the plan, need to skip this one though
			storeUpgrades := storetypes.StoreUpgrades{
				Added: []string{
					bech32ibctypes.ModuleName,
					oracletypes.ModuleName,
					leveragetypes.ModuleName,
				}, // We are adding these modules
				Renamed: nil,
				Deleted: nil,
			}

			// configure store loader that checks if version == upgradeHeight and applies
			// store upgrades
			app.SetStoreLoader(
				upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades),
			)
		}
	}

	// Calypso->Cosmwasm STORE LOADER SETUP
	// Register the new Calypso modules and the special StoreLoader to add them
	if upgradeInfo.Name == cosmwasm.PlanName {
		if !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) { // Recognized the plan, need to skip this one though
			storeUpgrades := storetypes.StoreUpgrades{
				Added: []string{
					wasm.ModuleName,
				}, // We are adding these modules
				Renamed: nil,
				Deleted: nil,
			}

			// configure store loader that checks if version == upgradeHeight and applies
			// store upgrades
			app.SetStoreLoader(
				upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades),
			)
		}
	}
}

// Name returns the name of the Umee network.
func (app *UmeeApp) Name() string {
	return app.BaseApp.Name()
}

// SimulationManager returns the application's SimulationManager.
func (app *UmeeApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// BeginBlocker implements Umee's BeginBlock ABCI method.
func (app *UmeeApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker implements Umee's EndBlock ABCI method.
func (app *UmeeApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer implements Umee's InitChain ABCI method.
func (app *UmeeApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(fmt.Sprintf("failed to unmarshal genesis state: %v", err))
	}

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	SetWasmDefaultGenesisState(app.appCodec, genesisState)

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height via Umee's BaseApp.
func (app *UmeeApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all of Umee's module account addresses.
func (app *UmeeApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns Umee's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *UmeeApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns Umee's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *UmeeApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Umee's InterfaceRegistry.
func (app *UmeeApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *UmeeApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *UmeeApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *UmeeApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *UmeeApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *UmeeApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// registry Tendermint RPC proxy routes
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)

	// register legacy tx routes
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)

	// register new tx routes from grpc-gateway
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register new tendermint queries routes from grpc-gateway
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register legacy and grpc-gateway routes for all modules
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *UmeeApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService
// method.
func (app *UmeeApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetBaseApp is used solely for testing purposes.
func (app *UmeeApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper is used solely for testing purposes.
func (app *UmeeApp) GetStakingKeeper() stakingkeeper.Keeper {
	return app.StakingKeeper
}

// GetIBCKeeper is used solely for testing purposes.
func (app *UmeeApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper is used solely for testing purposes.
func (app *UmeeApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig is used solely for testing purposes.
func (app *UmeeApp) GetTxConfig() client.TxConfig {
	return MakeEncodingConfig().TxConfig
}

// GetMaccPerms returns a deep copy of the module account permissions.
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}

func initParamsKeeper(
	appCodec codec.Codec,
	legacyAmino *codec.LegacyAmino,
	key, tkey sdk.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(gravitytypes.ModuleName)
	paramsKeeper.Subspace(leveragetypes.ModuleName)
	paramsKeeper.Subspace(oracletypes.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)

	return paramsKeeper
}

func getGovProposalHandlers() []govclient.ProposalHandler {
	proposalHandlers := []govclient.ProposalHandler{
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.ProposalHandler,
		upgradeclient.CancelProposalHandler,
		leverageclient.ProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
	}

	return append(proposalHandlers, wasmclient.ProposalHandlers...)
}

func VerifyAddressFormat(bz []byte) error {
	bzLen := len(bz)
	if bzLen == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "invalid address; cannot be empty")
	}

	if bzLen != AddrLen && bzLen != ContractAddrLen {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownAddress,
			"invalid address length; got: %d, should be: %d or %d for cosmwasm contract addr",
			bzLen, AddrLen, ContractAddrLen,
		)
	}

	return nil
}
