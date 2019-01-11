import * as React from 'react'
import * as StellarBase from "stellar-sdk"


// What we gonna use to access Horizon network
export interface StellarRepository {
    loadAccount(address: string): Promise<StellarBase.Account>

    transferAsset(
        destinationAddress: string,
        assetName: string,
        amount: string,
        issuer: string,
        sourceAddress: string,
        sourceSecret: string
    ): Promise<StellarBase.Server.TransactionRecord>
}

// Define default Horizon network access implementation
export class DefaultStellarRepository implements StellarRepository {

    constructor(server: StellarBase.Server) { }

    async loadAccount(address: string): Promise<StellarBase.Server.AccountResponse> {
        throw new Error("Implement me")
    }

    async transferAsset(destinationAddress: string,
                        assetName: string, amount: string,
                        issuer: string, sourceAddress: string,
                        sourceSecret: string
    ): Promise<StellarBase.Server.TransactionRecord> {

        throw new Error("Implement me")
    }

}


// Create context API to provider Horizon network access
// to our React Component and Hook
const TestNet = new StellarBase.Server('https://horizon-testnet.stellar.org')

export const HorizonContext = React.createContext(new DefaultStellarRepository(TestNet))
