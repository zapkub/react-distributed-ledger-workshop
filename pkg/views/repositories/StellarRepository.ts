import * as React from 'react'
import * as StellarBase from "stellar-sdk"
import {Network} from "stellar-sdk";


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

    private _server: StellarBase.Server;

    constructor(server: StellarBase.Server) {
        this._server = server
        Network.useTestNetwork()
    }

    async loadAccount(address: string): Promise<StellarBase.Server.AccountResponse> {
        return this._server.loadAccount(address)
    }


    async transferAsset(destinationAddress: string,
                        assetName: string, amount: string,
                        issuer: string, sourceAddress: string,
                        sourceSecret: string
    ): Promise<StellarBase.Server.TransactionRecord> {

        const source = await this._server.loadAccount(sourceAddress)
        const asset = new StellarBase.Asset(assetName, issuer)
        const tx = new StellarBase.TransactionBuilder(source)
            .addOperation(
                StellarBase.Operation.payment(
                    {
                        amount: amount,
                        asset: asset,
                        destination: destinationAddress,
                    }
                )
            )
            .build()
        const keypair = StellarBase.Keypair.fromSecret(sourceSecret)
        tx.sign(keypair)
        return this._server.submitTransaction(tx)
    }

}


// Create context API to provider Horizon network access
// to our React Component and Hook
const TestNet = new StellarBase.Server('https://horizon-testnet.stellar.org')

export const HorizonContext = React.createContext(new DefaultStellarRepository(TestNet))
