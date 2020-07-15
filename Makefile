generate-db-manifest:
	cat k8s/db/deployment.yaml | envsubst | tee .db_deployment.yaml

prepare: generate-db-manifest
	kubectl apply -f .db_deployment.yaml \
	&& echo "export DB_PORT=$$(kubectl get svc ci-with-skaffold-db -o jsonpath='{.spec.ports[0].nodePort}')"

cleanup: generate-db-manifest
	kubectl delete -f .db_deployment.yaml
