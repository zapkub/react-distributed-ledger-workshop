import * as React from 'react'
import Head from "next/head";

import * as StellarBase from "stellar-sdk"


const config = {
    "issuerAddress": "GCMTJ3FNBYPYNQFISRBC6DOFQTKWD7LZ6JUEPJPOQONB5TAQN6Z5XUOL",
    "assetName": "V4BP",
    "candidates": [
        {
            "name": "lisa",
            "address": "GCYVVFSLGH4AYKMVGLLSIYCWKMHPQHMZ3RHYCRGMD4BHNEHIGVNFQAAX"
        },
        {
            "name": "jennie",
            "address": "GBBWXRBPERMT5SBT7ZZMHOWGJS67DZQ5MFCKE5ZPUHZSEFDO5LCBYGRI"
        },
        {
            "name": "jisoo",
            "address": "GB6HDB7DWJOS7EWZH33OL2PLR6DW3NLA4JCRDXQFVEMU5ER5KAJNH2NX"
        },
        {
            "name": "rose",
            "address": "GDVVO5JNO5CXSNX7UTVEHCRXGO2ZYU7CDAAIUBU7VO35GFWSUQCXT7FU"
        }
    ]
}

const READY = 'ready'
const INVALID = 'invalid'
const VALID = 'valid'
function useAdressValidation() {

    const [address, setAddress] = React.useState<string>("")
    const [validateState, setValidateState] = React.useState(READY)

    return {
        address,
        setAddress: (address) => {
            setValidateState(READY)
           setAddress(address)
        },

        validateState,
        doValidate: async () => {

            const TestNet = new StellarBase.Server('https://horizon-testnet.stellar.org')

            const account = await TestNet.loadAccount(address)
            const voteBalance = account.balances.find(
                (balance) => {
                    if(balance.asset_type == "credit_alphanum4") {
                        return balance.asset_code === config.assetName &&
                            balance.asset_issuer === config.issuerAddress
                    }
                    return false
                }
            )

            if (voteBalance) {
                console.log(voteBalance)
                setValidateState(VALID)
            } else {
                console.error("NO asset")
                setValidateState(INVALID)
            }


        }
    }
}

export default () => {

    const validationState = useAdressValidation()

    const address = validationState.address
    const setAddress = validationState.setAddress
    const doValidate = validationState.doValidate
    const inputState = validationState.validateState

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
                    {`Blackpink General Election 2019 ${config.assetName}`}
                    <div className="input-group mb-3">
                        <div className="input-group-prepend">
                            <span className="input-group-text">{"Public key"}</span>
                        </div>

                        <input
                            value={address}
                            onChange={
                                (e) => {
                                    setAddress(e.target.value)
                                }
                            }


                            type="text" className={`form-control ${ inputState === INVALID ? "is-invalid" : ""  }`}
                            placeholder="publickey"
                        />
                        <div className="invalid-feedback">
                            Invalid account, no V4BP asset
                        </div>


                    </div>

                    <button onClick={doValidate} className={"btn btn-primary"}>{"Validate address"}</button>


                </div>
            </div>
        </div>
    )


}