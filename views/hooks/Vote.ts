import * as React from 'react'
import {HorizonContext} from "../repositories/StellarRepository";
import {ServiceContext} from "../repositories/ServiceRepository";
import {Server} from "stellar-sdk";

export function useVote(address: string, callback: (tx: Server.TransactionRecord) => void) {

    const [secret, setSecret] = React.useState("")
    const [amount, setAmount] = React.useState(1)
    const [candidateAddress, setCandidateAddress] = React.useState<string | undefined>(undefined)

    const horizon = React.useContext(HorizonContext);
    const service = React.useContext(ServiceContext)

    return {

        secret,
        setSecret,
        amount,
        setAmount,
        setCandidateAddress,
        candidateAddress,
        voteForCandidateAddress: async () => {

            if (candidateAddress) {
                console.log(secret)
                const config = await service.getConfiguration()
                try {
                    const txResult = await horizon.transferAsset(
                        candidateAddress,
                        config.assetName,
                        "1",
                        config.issuerAddress,
                        address,
                        secret,
                    )
                    callback(txResult)

                }catch(e) {
                    console.error(e)
                }

            }

        }

    }
}