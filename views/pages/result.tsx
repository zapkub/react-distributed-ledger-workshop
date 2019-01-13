import * as React from 'react'
import {useConfiguration} from "../hooks/Configuration";
import {useElectionResult} from "../hooks/ElectionResult";


export default () => {

    const {isLoading, configuration} = useConfiguration()
    if (isLoading || !configuration) {
        return (
            <div>
                {"Preparing..."}
            </div>
        )
    }
    const {results,  isLoading: electionLoading} = useElectionResult(configuration)

    if (electionLoading) {
        return (
            <div>
                {"fetching election result..."}
            </div>
        )
    }
    return (
        <div className={"container mx-3 my-3"}>
            <h1>{'Election result'}</h1>
            {results.map(
                candidate => {
                    return (
                        <div className={"my-1"} key={candidate.name}>
                            {`${candidate.name}: ${candidate.amount}`}
                        </div>
                    )

                }
            )}
        </div>
    )

}