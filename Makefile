.PHONY: all

project_id := otomarukanta-a
version := 1

serve:
	goapp serve app

deploy:
	goapp deploy -application ${project_id} -version ${version} app

