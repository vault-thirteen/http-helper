////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
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
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/vault-thirteen/errorz"
)

// Functions which help in receiving Data from HTTP Requests.

// ReceiveJSON Function receives an Object encoded with JSON Format from the
// HTTP Request's Body.
func ReceiveJSON(
	r *http.Request,
	receiver interface{},
) (err error) {

	var bodyContents []byte
	var jsonDecoder *json.Decoder

	if reflect.TypeOf(receiver).Kind() != reflect.Ptr {
		return errors.New(ErrNotPointer)
	}

	bodyContents, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func() {
		var derr error
		derr = r.Body.Close()
		err = errorz.Combine(err, derr)
	}()

	jsonDecoder = json.NewDecoder(bytes.NewReader(bodyContents))
	err = jsonDecoder.Decode(receiver)
	if err != nil {
		return err
	}

	return nil
}
