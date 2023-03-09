# Chain

## Source Code

---

protobuf: [git@github.com](mailto:git@github.com):Raven0520/Proto.git

rpc service: [git@github.com](mailto:git@github.com):Raven0520/BTC.git

http: [git@github.com](mailto:git@github.com):Raven0520/Wallet.git

### Questions

- What is the best way(s) to provide the seed onto this server? Please justify the approach(es) with reasons.
    
    The best way for users to provide seed to the platform is to only provide seed of non-main wallets, because in the blockchain network, any behavior of exposing mnemas, seed and private keys is unsafe. It is more reasonable to only use non-main wallets to interact with the platform, and then transfer assets to the main wallet addresses. At the same time, users can also use multi-signature technology to protect assets in non-primary wallets.
    

### Requirements

- Is the generated address correct?
    
    Importing the wallet address into the appropriate blockchain network can verify whether the address is correct. At present, only manual import verification is carried out, and it is directly implemented without unit testing.
    
- Is it safe for users to use?
    
    Not very safe for now. Users should use non-primary wallets.
    
- Is the documentation easy to read?
    
    Because the function is not complicated, the document is also very simple.
    
- Does it follow best practices?
    
    I think so at present.
    
- Add features you can think of on top of the basic requirements?
    
    The requirements document does not further elaborate on the usage scenario after the generated address. Assuming that the generated address is only used in the platform, BTC PRC service can add the generated rules to ensure that the third party cannot generate the same address and private key after the mnie is stolen.
    
    ### Improve security
    
    - Modify the generation rules
    - Map the mnemonics provided by users. This will cause users to use the same mnemonic and path cannot correctly export the wallet address to other platforms.