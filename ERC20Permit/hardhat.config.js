require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
  solidity: "0.8.20",
  networks: {
    mainnet: {
      url: process.env.MAINNET_RPC,
      accounts: [process.env.PRIVATE_KEY],
      gasPrice: 30_000_000_000, // 30 Gwei，可根据网络情况调整
    },
    hardhat: {},
  },
};
