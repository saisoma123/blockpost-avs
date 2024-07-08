// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {ECDSAStakeRegistry} from "@eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import {Quorum, StrategyParams} from "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistryEventsAndErrors.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {HelloWorldServiceManager, IServiceManager} from "../src/HelloWorldServiceManager.sol";
import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract BlockPostDeployer is Script, Utils {
    // BlockPost contracts
    ProxyAdmin public blockPostProxyAdmin;
    PauserRegistry public blockPostPauserReg;

    ECDSAStakeRegistry public stakeRegistryProxy;
    ECDSAStakeRegistry public stakeRegistryImplementation;

    BlockPost public blockPostServiceManagerProxy;
    BlockPost public blockPostServiceManagerImplementation;

    function run() external {
        // EigenLayer contracts
        //address strategyManagerAddress = 0x...; //
        //address delegationManagerAddress = 0x...; //
        //address avsDirectoryAddress = 0x...; //
        //address eigenLayerProxyAdminAddress = ; //
        //address eigenLayerPauserRegAddress = ; //

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

        address blockPostCommunityMultisig = msg.sender;
        address blockPostPauser = msg.sender;

        vm.startBroadcast();
        _deployBlockPostContracts(
            delegationManager,
            avsDirectory,
            blockPostCommunityMultisig,
            blockPostPauser
        );
        vm.stopBroadcast();
    }

    function _deployBlockPostContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        address blockPostCommunityMultisig,
        address blockPostPauser
    ) internal {
        blockPostProxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](2);
        pausers[0] = blockPostPauser;
        pausers[1] = blockPostCommunityMultisig;
        blockPostPauserReg = new PauserRegistry(
            pausers,
            blockPostCommunityMultisig
        );

        EmptyContract emptyContract = new EmptyContract();

        blockPostServiceManagerProxy = BlockPostServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(blockPostProxyAdmin),
                    ""
                )
            )
        );
        stakeRegistryProxy = ECDSAStakeRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(blockPostProxyAdmin),
                    ""
                )
            )
        );
    }
}
