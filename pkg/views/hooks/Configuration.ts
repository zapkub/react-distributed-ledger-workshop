import * as React from 'react'
import {GetConfigurationResponseBody, ServiceContext} from "../repositories/ServiceRepository";

export function useConfiguration() {
    const [isLoading, setIsLoading] = React.useState(true)
    const [configuration, setConfiguration] = React.useState<GetConfigurationResponseBody | undefined >(undefined)

    const service = React.useContext(ServiceContext)

    async function refetch() {
        setIsLoading(true)
        const result = await service.getConfiguration()
        setConfiguration(result)
        setIsLoading(false)
    }

    React.useEffect(() => {
        refetch()
    }, [])

    return {
        isLoading,
        configuration: configuration,
        refetch
    }

}


