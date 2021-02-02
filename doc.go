/*
Copyright 2017 Google Inc. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package tika provides a client for using Apache Tika's (http://tika.apache.org) Server API.

To parse the contents of a file (or any io.Reader), you will need to open the io.Reader,
create a client, and call client.Parse.

	// import "os"
	f, err := os.Open("path/to/file")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	serverURL := "TODO"

	client := tika.NewClient(nil, serverURL)
	body, err := client.Parse(context.Background(), f)

If you pass an *http.Client to tika.NewClient, it will be used for all requests.

Some functions return a custom type, like Parsers(), Detectors(), and
MIMETypes(). Use these to see what features are supported by the current
Tika server.
*/
package tika
