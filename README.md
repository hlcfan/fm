# fm

Format JSON, XML string. I mainly use it in Vim buffers, so that I can format
any JSON/XML string without knowing type.

Before:
```
# format JSON
:%! jq .

# format XML
:%! xmllint --format -
```

After:
```
:%! fm
```

### Usage

```
echo "{\"key\":\"value\"}" | fm
{
  "key": "value"
}
```

### Note

If on MacOS, you'll need to codesign the binary file.
```
codesign -s - fm
```
