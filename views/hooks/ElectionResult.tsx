import * as React from 'react'
import {GetConfigurationResponseBody} from "../repositories/ServiceRepository";
import {HorizonContext} from "../repositories/StellarRepository";

interface CandidateResultItem {
    name: string
    amount: number
}

export function useElectionResult(configuration: GetConfigurationResponseBody) {

    const [results, setResults] = React.useState<CandidateResultItem[]>([])
    const [isLoading, setIsLoading] = React.useState(true)
    const horizon = React.useContext(HorizonContext)

    React.useEffect(() => {
        setIsLoading(true)
        const candidateBalancePromises = configuration.candidates.map<Promise<CandidateResultItem>>(async candidate => {
            const account = await horizon.loadAccount(candidate.address)
            const voteAsset = account.balances.find((asset) => {
                if (asset.asset_type == "credit_alphanum4") {
                    return asset.asset_issuer == configuration.issuerAddress && asset.asset_code == configuration.assetName
                } else {
                    return false
                }
            })

            if (voteAsset) {
                return {
                    amount: parseInt(voteAsset.balance),
                    name: candidate.name,
                }
            } else {
                throw new Error(`${candidate.name} asset trust error`)
            }
        })
        Promise.all(candidateBalancePromises).then((resp) => {
            setResults(resp)
            setIsLoading(false)
        })
    }, [])


    return {
        results,
        isLoading,
    }

}