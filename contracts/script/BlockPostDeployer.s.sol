// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {ECDSAStakeRegistry} from "eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import {Quorum, StrategyParams} from "eigenlayer-middleware/src/interfaces/IECDSAStakeRegistryEventsAndErrors.sol";
import "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {BlockPostServiceManager} from "../src/BlockPostServiceManager.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";
import "./SimpleProxy.sol";

contract BlockPostDeployer is Script {
    // BlockPost contracts
    ProxyAdmin public blockPostProxyAdmin;
    PauserRegistry public blockPostPauserReg;

    SimpleProxy public stakeRegistryProxy;
    ECDSAStakeRegistry public stakeRegistryImplementation;

    SimpleProxy public blockPostServiceManagerProxy;
    BlockPostServiceManager public blockPostServiceManagerImplementation;

    event Initialized(address value);

    function run() external {
        // EigenLayer contracts
        address strategyManagerAddress = 0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6;
        address delegationManagerAddress = 0xA44151489861Fe9e3055d95adC98FbD462B948e7;
        address avsDirectoryAddress = 0x055733000064333CaDDbC92763c58BF0192fFeBf;
        address eigenLayerProxyAdminAddress = 0xDB023566064246399b4AE851197a97729C93A6cf;
        address eigenLayerPauserRegAddress = 0x85Ef7299F8311B25642679edBF02B62FA2212F06;
        address baseStrategyImplementationAddr = 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9;

        IStrategyManager strategyManager = IStrategyManager(
            strategyManagerAddress
        );
        IDelegationManager delegationManager = IDelegationManager(
            delegationManagerAddress
        );
        IAVSDirectory avsDirectory = IAVSDirectory(avsDirectoryAddress);
        ProxyAdmin eigenLayerProxyAdmin = ProxyAdmin(
            eigenLayerProxyAdminAddress
        );
        PauserRegistry eigenLayerPauserReg = PauserRegistry(
            eigenLayerPauserRegAddress
        );
        StrategyBase baseStrategyImplementation = StrategyBase(
            baseStrategyImplementationAddr
        );

        address blockPostCommunityMultisig = msg.sender;
        address blockPostPauser = msg.sender;

        vm.startBroadcast();
        _deployBlockPostContracts(
            delegationManager,
            avsDirectory,
            baseStrategyImplementation,
            blockPostCommunityMultisig,
            blockPostPauser
        );
        vm.stopBroadcast();
    }

    function _deployBlockPostContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        IStrategy baseStrategyImplementation,
        address blockPostCommunityMultisig,
        address blockPostPauser
    ) internal {
        // Initializes Proxy Admin, Service Manager proxy, and Empty Contract

        blockPostProxyAdmin = new ProxyAdmin(msg.sender);

        {
            address[] memory pausers = new address[](2);
            pausers[0] = blockPostPauser;
            pausers[1] = blockPostCommunityMultisig;
            blockPostPauserReg = new PauserRegistry(
                pausers,
                blockPostCommunityMultisig
            );
        }

        EmptyContract emptyContract = new EmptyContract();

        blockPostServiceManagerProxy = new SimpleProxy(address(emptyContract));

        // Initializes SimpleProxy and sets to empty contract initially
        stakeRegistryProxy = stakeRegistryProxy = new SimpleProxy(
            address(emptyContract)
        );

        // Upgrades the proxy to the stake registry implementation
        {
            stakeRegistryImplementation = new ECDSAStakeRegistry(
                delegationManager
            );

            SimpleProxy(payable(address(stakeRegistryProxy))).upgradeTo(
                address(stakeRegistryImplementation)
            );
        }

        // Initializes a quorum and initializes stake registry with service manager and
        // quorum
        {
            StrategyParams memory strategyParams = StrategyParams({
                strategy: baseStrategyImplementation,
                multiplier: 10_000
            });

            StrategyParams[]
                memory quorumsStrategyParams = new StrategyParams[](1);
            quorumsStrategyParams[0] = strategyParams;

            Quorum memory quorum = Quorum(quorumsStrategyParams);

            stakeRegistryImplementation.initialize(
                address(blockPostServiceManagerProxy),
                1,
                quorum
            );
        }

        // Makes the new service manager with new initialized proxies
        blockPostServiceManagerImplementation = new BlockPostServiceManager(
            avsDirectory,
            IRegistryCoordinator(address(0)),
            IStakeRegistry(address(stakeRegistryProxy))
        );

        // Upgrades the SimpleProxy to point to service manager implementation
        SimpleProxy(payable(address(blockPostServiceManagerProxy))).upgradeTo(
            address(blockPostServiceManagerImplementation)
        );

        console.log(
            "BlockPostServiceManager Proxy:",
            address(blockPostServiceManagerProxy)
        );
        console.log(
            "BlockPostServiceManager Implementation:",
            address(blockPostServiceManagerImplementation)
        );
        console.log("StakeRegistry Proxy:", address(stakeRegistryProxy));
        console.log(
            "StakeRegistry Implementation:",
            address(stakeRegistryImplementation)
        );
    }
}
