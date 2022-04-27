build:
	go build -o ruby/fluent-plugin-encrypt/lib/fluent/plugin/go/encryptor.so -buildmode=c-shared encryptor.go
	gem build fluent-plugin-encrypt.gemspec -o fluentbit-plugin-encrypt.gem -C ruby/fluent-plugin-encrypt

install:
	/usr/bin/gem install fluentbit-plugin-encrypt.gem

start:
	/usr/local/bin/fluentd -c ./fluent.example.conf -vv