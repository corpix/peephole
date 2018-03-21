config.json: config.nix config.mk
	nix-instantiate                                \
		--eval                                 \
		--expr "builtins.toJSON (import ./$<)" \
	| jq -r . | jq -S . > $@
