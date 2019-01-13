import * as React from 'react'
import Head from "next/head";
import Router from "next/router"
import {AddressValidateStatus, useAddressValidation} from "../hooks/AddressValidation";
import {useVote} from "../hooks/Vote";
import {useConfiguration} from "../hooks/Configuration";


export default () => {

    const {isLoading, configuration} = useConfiguration()

    const {
        address,
        setAddress,
        doValidateAddress,
        isValidate,
        balance
    } = useAddressValidation()

    const {
        setSecret,
        secret,
        candidateAddress,
        setCandidateAddress,
        setAmount,
        amount,
        voteForCandidateAddress
    } = useVote(address, (tx) => {
       alert(tx.hash)
        Router.push("/result")
    })

    if (isLoading) {
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

                <div>
                    <div className={"h2"}>
                        {"Account information"}
                        <div className={"input-group my-3"}>
                            <div className="input-group-prepend">
                                <span className="input-group-text" id="basic-addon1">{"Address"}</span>
                            </div>
                            <input type={"text"}
                                   value={address}
                                   onChange={(e) => setAddress(e.target.value)}
                                   className={`form-control ${isValidate === AddressValidateStatus.INVALID ? "is-invalid" : ""}`}/>
                            <div className="invalid-feedback">
                                {"Please provide a valid address."}
                            </div>
                        </div>

                    </div>
                </div>

                <button className={"btn btn-primary"} onClick={doValidateAddress}>{"Validate"}</button>

            </div>
            {
                isValidate === AddressValidateStatus.VALID ? (
                    <div className="container mx-3">
                        <div>{'Vote form'}</div>

                        <div className={"input-group my-3"}>
                            <div className="input-group-prepend">
                                <span className="input-group-text" id="basic-addon1">{"Vote credit"}</span>
                            </div>
                            <input disabled={true}
                                   value={balance}
                                   className={"form-control"}
                            />
                        </div>

                        <div className={"input-group my-3"}>
                            <div className="input-group-prepend">
                                <span className="input-group-text" id="basic-addon1">{"Vote amount"}</span>
                            </div>
                            <input
                                onChange={(e) => setAmount(parseInt(e.target.value) || 0)}
                                value={amount}
                                type={"number"}
                                className={"form-control"}
                            />
                        </div>

                        {
                            configuration ? configuration.candidates.map((candidate) => {
                                return (
                                    <div key={candidate.name} className={"input-group my-3"}>
                                        <button
                                            onClick={() => setCandidateAddress(candidate.address)}
                                            className={`btn btn-secondary ${candidateAddress === candidate.address ? "btn-success" : ""}`}>{"Vote for " + candidate.name}</button>
                                    </div>
                                )
                            }) : null
                        }

                        <div className={"input-group my-3"}>
                            <div className="input-group-prepend">
                                <span className="input-group-text" id="basic-addon1">{"Secret"}</span>
                            </div>
                            <input
                                    onChange={(e) => setSecret(e.target.value)}
                                   value={secret} type={"password"}
                                   className={"form-control"}/>
                        </div>

                        <div className={"btn btn-primary"} onClick={voteForCandidateAddress}>{'Confirm'}</div>
                    </div>
                ) : null
            }
        </div>
    )


}