package types

import (
	"fmt"
	"strings"
)

var CustomAccountParser = []string{ // for desmos
	"ownerAddress", "creator", "toAddress", "granter", "grantee", "owner", "withdraw_address",
}

var DefaultAccountParser = []string{
	"signer", "sender", "to_address", "from_address", "delegator_address",
	"validator_address", "submitter", "proposer", "depositor", "voter",
	"validator_dst_address", "validator_src_address",
}

func MessageParser(msg map[string]interface{}) (addresses string) {
	accountParser := append(DefaultAccountParser, CustomAccountParser...)

	addresses += "{"
	for _, role := range accountParser {
		if address, ok := msg[role].(string); ok {
			addresses += address + ","
		}
	}

	msgType := msg["@type"].(string)[1:]

	if msgType == "firmachain.firmachain.contract.MsgCreateContractFile" {

		msgText := fmt.Sprint(msg["ownerList"])

		fmt.Println(msgText)

		totalLength := len(msgText)

		fmt.Println("totalLength")
		fmt.Println(totalLength)

		if totalLength > 0 {

			for i := 0; i < totalLength; i++ {
				idx := strings.Index(msgText, "firma1")
				if idx != -1 {
					const lenghOfAddress = 44
					temp := msgText[idx : idx+lenghOfAddress]
					i += (idx + lenghOfAddress)

					if !strings.Contains(addresses, temp) {
						addresses += temp + ","
					}

					fmt.Println(temp)
				}
			}

			for i := 0; i < totalLength; i++ {
				idx := strings.Index(msgText, "firmavaloper1")
				if idx != -1 {
					const lenghOfValidatorAddress = 51
					temp := msgText[idx : idx+lenghOfValidatorAddress]
					i += (idx + lenghOfValidatorAddress)

					if !strings.Contains(addresses, temp) {
						addresses += temp + ","
					}
					fmt.Println(temp)
				}
			}
		}
	}

	if msgType == "cosmos.authz.v1beta1.MsgExec" {

		msgText := fmt.Sprint(msg["msgs"])
		fmt.Println(msgText)

		totalLength := len(msgText)

		fmt.Println("totalLength")
		fmt.Println(totalLength)

		if totalLength > 0 {

			for i := 0; i < totalLength; i++ {
				idx := strings.Index(msgText, "firma1")
				if idx != -1 {
					const lenghOfAddress = 44
					temp := msgText[idx : idx+lenghOfAddress]
					i += (idx + lenghOfAddress)

					if !strings.Contains(addresses, temp) {
						addresses += temp + ","
					}

					fmt.Println(temp)
				}
			}

			for i := 0; i < totalLength; i++ {
				idx := strings.Index(msgText, "firmavaloper1")
				if idx != -1 {
					const lenghOfValidatorAddress = 51
					temp := msgText[idx : idx+lenghOfValidatorAddress]
					i += (idx + lenghOfValidatorAddress)

					if !strings.Contains(addresses, temp) {
						addresses += temp + ","
					}
					fmt.Println(temp)
				}
			}
		}

		// fmt.Println(msgType)
		// fmt.Println(addresses)
		// os.Exit(3)
	}

	if input, ok := msg["input"].([]map[string]interface{}); ok {
		for _, i := range input {
			addresses += i["address"].(string) + ","
		}
	}

	if output, ok := msg["output"].([]map[string]interface{}); ok {
		for _, i := range output {
			addresses += i["address"].(string) + ","
		}
	}

	if len(addresses) == 1 {
		return "{}"
	}

	addresses = addresses[:len(addresses)-1] // remove trailing ,
	addresses += "}"

	return addresses
}
