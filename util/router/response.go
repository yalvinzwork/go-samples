package router

import ()

// func WriteResponse(ResultJsonString map[string]string, r *Request, w http.ResponseWriter) error {

// 	resp := WSResponse{
// 		Status: "OK",
// 		//Config: "null",
// 		Data: ResultJsonString,
// 	}

// 	resp.ServerProcessTime = fmt.Sprintf("%v", time.Since(r.start).Seconds())
// 	json, err := json.Marshal(resp)
// 	if err != nil {
// 		return err
// 	}

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(json))

// 	return nil
// }
