INIT_OPTS=init \
	--workload-config .operator-builder/workload.yaml \
   	--repo github.com/lander2k2/pocket-v1-operator \
	--controller-image lander2k2/pocket-v1-operator:latest \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .operator-builder/workload.yaml \
	--controller \
	--resource# \
	#--force

operator-clean:
	rm -rf ../*

operator-init:
	cd .. && operator-builder $(INIT_OPTS)

operator-create:
	cd .. && operator-builder $(CREATE_OPTS)

operator-dependencies:
	kubectl apply -f dependencies

operator-samples:
	cp samples/* ../config/samples/

