const { ethers } = require("hardhat");

async function main() {
  const name = "MyPermitToken";
  const symbol = "MPT";

  const ERC20Permit = await ethers.getContractFactory("ERC20Permit");
  const token = await ERC20Permit.deploy(name, symbol);

  await token.deployed();

  console.log(`ERC20Permit deployed to: ${token.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Error deploying:", error);
    process.exit(1);
  });
