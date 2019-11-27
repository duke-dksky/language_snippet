
[参考文章1](https://stackoverflow.com/questions/21197239/decoding-json-using-json-unmarshal-vs-json-newdecoder-decode)
[参考文章2](https://sanyuesha.com/2018/05/07/go-json/)

编码是指编成json格式,解码是指从json格式解析
Decoding JSON using json.Unmarshal vs json.NewDecoder.Decode

* Marshal and Unmarshal convert JSON into a string and vice versa. 
  - Marshal: This will marshal the JSON into []bytes 
  ```
  func Marshal(v interface{}) ([]byte, error)
  ```
  - Unmarshal: This will unmarshal the JSON from []bytes
  ```
  func Unmarshal(data []byte, v interface{}) error
  ```

* Encoding and decoding convert JSON into a stream and vice versa
  - Encoding: write struct to slice of a stream (bytes.buffer)
  - Decoding: read data from a slice of a stream and convert it into a struct (bytes.buffer)

The only difference is if you wanna play with string or bytes use marshal 
and if any data you want to read or write to some writer interface, use encodes and decode.

