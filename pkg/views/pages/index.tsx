import * as React from 'react'
import Head from "next/head";
import {useConfiguration} from "../hooks/Configuration";


export default () => {

    const {isLoading, configuration} = useConfiguration()

    if (isLoading || !configuration) {
        return (
            <div>
                {"Preparing..."}
            </div>
        )
    }


    return (
        <div>
            <Head>
                <link rel="stylesheet"
                      href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css"
                      integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO"
                      crossOrigin="anonymous"/>
            </Head>

            <div className={"container mx-3 my-3"}>
                <div className={"h1"}>
                    {"Blackpink General Election 2019"}
                </div>
            </div>
        </div>
    )


}