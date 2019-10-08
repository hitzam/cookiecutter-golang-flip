module github.com/kitabisa/{{ cookiecutter.app_name }}

require (
	{%- if cookiecutter.use_logrus_logging == "y" %}
    github.com/sirupsen/logrus v1.4.2
    {%- endif %}
    {%- if cookiecutter.use_viper_config == "y" %}
    github.com/spf13/viper v1.4.0
    {%- endif %}
	{%- if cookiecutter.use_cobra_cmd == "y" %}
    github.com/spf13/cobra v0.0.5
    {%- endif %}
	{%- if cookiecutter.use_migrate_migration == "y" %}
    github.com/golang-migrate/migrate v4.6.2
    {%- endif %}
)
