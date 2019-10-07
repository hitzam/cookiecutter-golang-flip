FROM kitabisa/debian-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/kitabisa/{{ cookiecutter.app_name }}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/{{ cookiecutter.app_name }}/bin

WORKDIR /opt/{{ cookiecutter.app_name }}/bin

COPY ./bin/{{ cookiecutter.app_name }} /opt/{{ cookiecutter.app_name }}/bin/
{% if cookiecutter.use_migrate_migration == "y" -%}
COPY ./migrations /opt/{{ cookiecutter.app_name }}/bin/migrations/
{% endif %}
RUN chmod +x /opt/{{ cookiecutter.app_name }}/bin/{{ cookiecutter.app_name }}

# Create appuser
RUN adduser --disabled-password --gecos '' {{ cookiecutter.app_name }}
USER {{ cookiecutter.app_name }}

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/{{ cookiecutter.app_name }}/bin/{{ cookiecutter.app_name }}"]
