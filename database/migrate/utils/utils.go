package types

import (
	"fmt"
	"os"
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

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}

	length := len(value)

	posLast := strings.Index(value[posFirst:length], b)
	if posLast == -1 {
		return ""
	}
	return (value[posFirst:length])[:posLast]
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

		fmt.Println(msgType)
		fmt.Println(addresses)

		msgText := fmt.Sprint(msg["ownerList"])

		if len(msgText) > 0 {

			trimmedStr := strings.Trim("msgText", "[]")
			strList := strings.Split(trimmedStr, " ")

			for _, str := range strList {

				if len(str) > 0 {
					addresses += str + ","
				}
			}
		}

		fmt.Println(msgType)
		fmt.Println(addresses)
		os.Exit(3)
	}

	if msgText, ok := msg["msgs"].(string); ok {

		slice := strings.Split(msgText, " ")

		for _, str := range slice {

			userAddress := between(str, "firma1", "\"")

			if len(userAddress) > 0 {
				addresses += userAddress + ","
			}

			valoperAddress := between(str, "firmavaloper1", "\"")

			if len(valoperAddress) > 0 {
				addresses += userAddress + ","
			}
		}

		fmt.Println("msgs")
		os.Exit(3)
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
