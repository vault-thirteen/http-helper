// reply.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package httphelper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/vault-thirteen/header"
	"github.com/vault-thirteen/mime"
)

// Errors.
const (
	ErrNullPointer = "Null Pointer"
)

// Functions which help in replying to HTTP Requests.

// ReplyTextWithCode Function replies to the HTTP Request with the specified
// Text and HTTP Status Code.
func ReplyTextWithCode(
	w http.ResponseWriter,
	httpStatusCode int,
	replyText string,
) {
	var xerr error

	w.WriteHeader(httpStatusCode)
	_, xerr = w.Write([]byte(replyText))
	if xerr != nil {
		log.Println(replyText)
		log.Println(xerr)
	}
}

// ReplyErrorWithCode Function replies to the HTTP Request with an Error and
// the specified HTTP Status Code.
func ReplyErrorWithCode(
	w http.ResponseWriter,
	httpStatusCode int,
	err error,
) {
	ReplyTextWithCode(w, httpStatusCode, err.Error())
}

// ReplyErrorInternal Function replies to the HTTP Request with an Error and
// 'Internal Server Error' HTTP Status Code.
func ReplyErrorInternal(
	w http.ResponseWriter,
	err error,
) {
	ReplyErrorWithCode(w, http.StatusInternalServerError, err)
}

// ReplyJSON Function sends an Object in JSON Format to the HTTP Output Stream.
func ReplyJSON(
	w http.ResponseWriter,
	replyObject interface{},
) {
	var err error
	var response []byte

	// Encode an Object with JSON Format.
	response, err = json.Marshal(replyObject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	// Send the Reply.
	w.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

// ReplyJSONFast Function sends an Object in JSON Format to the HTTP Output
// Stream using the faster but less secure Way than an ordinary 'ReplyJSON'
// Method.
func ReplyJSONFast(
	w http.ResponseWriter,
	replyObject interface{},
) {
	var err error
	var jsonEncoder *json.Encoder

	// Create the JSON Encoder.
	jsonEncoder = json.NewEncoder(w)
	if jsonEncoder == nil {
		err = errors.New(ErrNullPointer)
		log.Println(err)
	}

	// Encode an Object with JSON Format and send it simultaneously.
	w.Header().Set(header.HttpHeaderContentType, mime.TypeApplicationJson)
	err = jsonEncoder.Encode(replyObject)
	if err != nil {
		log.Println(err)
	}
}
