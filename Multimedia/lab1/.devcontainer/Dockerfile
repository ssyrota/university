FROM condaforge/miniforge3:latest
USER root

COPY ./environment.yml /requirements/

RUN mamba env create --file /requirements/environment.yml
RUN mamba env list

RUN mamba install ipython_genutils
RUN mamba install ipykernel
RUN ipython kernel install --user --name=env1

ENTRYPOINT ["tail", "-f", "/dev/null"]