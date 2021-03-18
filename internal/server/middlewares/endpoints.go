/*
 * Copyright (C) 2020  SuperGreenLab <towelie@supergreenlab.com>
 * Author: Constantin Clauzel <constantin.clauzel@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// OutputObjectID - returns the inserted object ID
func OutputObjectID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := r.Context().Value(InsertedIDContextKey{}).(uuid.UUID)
	response := struct {
		ID uuid.UUID `json:"id"`
	}{id}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("json.NewEncoder in OutputObjectID %q - %+v", err, response)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// OutputObjectID - returns the inserted object ID
func OutputMultipleObjectIDs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ids := r.Context().Value(MultipleInsertedIDsContextKey{}).([]uuid.UUID)
	err := r.Context().Value(MultipleInsertErrorContextKey{})

	errMsg := ""
	if e, ok := err.(error); ok {
		errMsg = e.Error()
	}
	response := struct {
		IDs []uuid.UUID `json:"ids"`
		Err string      `json:"error,omitempty"`
	}{ids, errMsg}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf("json.NewEncoder in OutputMultipleObjectIDs %q - %+v", err, response)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
