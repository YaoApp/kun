# GRPC Plugin

## Updating the Protocol

If you update the protocol buffers file, you can regenerate the file using the following command from this directory. You do not need to run this if you're just trying the example.

```bash
protoc -I proto/ proto/model.proto  --go_out=plugins=grpc:proto/
```

## More

https://github.com/hashicorp/go-plugin/tree/master/examples/grpc
