/*
 * Copyright (c) 2015 GRNET S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the
 * License. You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an "AS
 * IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language
 * governing permissions and limitations under the License.
 *
 * The views and conclusions contained in the software and
 * documentation are those of the authors and should not be
 * interpreted as representing official policies, either expressed
 * or implied, of GRNET S.A.
 *
 */

package statusServices

import "encoding/xml"

// InputParams struct holds as input all the url params of the request
type InputParams struct {
	startTime string // UTC time in W3C format
	endTime   string
	report    string
	groupType string
	group     string
	service   string
	format    string
}

// DataOutput struct holds the queried data from datastore
type DataOutput struct {
	Report        string `bson:"report"`
	Timestamp     string `bson:"timestamp"`
	EndpointGroup string `bson:"endpoint_group"`
	Service       string `bson:"service"`
	Status        string `bson:"status"`
	DateInteger   string `bson:"date_int"`
}

// xml/json response related structs

type rootOUT struct {
	XMLName        xml.Name            `xml:"root" json:"-"`
	EndpointGroups []*endpointGroupOUT `json:"groups"`
}

type endpointGroupOUT struct {
	XMLName   xml.Name      `xml:"group" json:"-"`
	Name      string        `xml:"name,attr" json:"name"`
	GroupType string        `xml:"type,attr" json:"type"`
	Services  []*serviceOUT `json:"services"`
}

type serviceOUT struct {
	XMLName   xml.Name     `xml:"group" json:"-"`
	Name      string       `xml:"name,attr" json:"name"`
	GroupType string       `xml:"type,attr" json:"type"`
	Statuses  []*statusOUT `json:"statuses"`
}

type statusOUT struct {
	XMLName   xml.Name `xml:"status" json:"-"`
	Timestamp string   `xml:"timestamp,attr" json:"timestamp"`
	Value     string   `xml:"value,attr" json:"value"`
}

// Message struct to hold the xml/json response
type messageOUT struct {
	XMLName xml.Name `xml:"root" json:"-"`
	Message string   `xml:"message" json:"message"`
}
