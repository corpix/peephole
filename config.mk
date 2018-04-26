config.json: config.nix config.mk
	nix-instantiate                          \
		--eval                                 \
		--expr "builtins.toJSON (import ./$<)" \
	| jq -r . | jq -S . > $@

config.yaml: config.json config.mk
	cat $< | formats --from json --to yaml > $@

.PHONY: config
config: config.json config.yaml
