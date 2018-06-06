# Key Value Playground

Simple WebSocket server that receive JSON payloads with keys and reply responses
(values) from a Postgres DB.

## Dictionary

The dictionary have only ten keys, from `0` to `9`, and each key has the value
as the translation of the key on a popular idiom.

| Key | Value |
|-----|---|
| '0' | '零' |
| '1' | 'Uno' |
| '2' | 'Two' |
| '3' | 'तीन' |
| '4' | 'أربعة' |
| '5' | 'Cinco' |
| '6' | 'ছয়' |
| '7' | 'семь' |
| '8' | '八' |
| '9' | 'ਨੌਂ' |

### Sample request and response payloads

Request:

```js
{
  key: '0'
}
```

Response:

```js
{
  value: '零'
}
```

## Initial release

For the initial release, the following tasks are expected:

* [x] Setup a relation database
* [x] Create the table `keys` (id, key, value) with the above dictionary
* [ ] Implement a WebSocket server to handle JSON request/responses
* [ ] Implement automated tests to check all dictionary key/values
* [ ] Deploy the server
