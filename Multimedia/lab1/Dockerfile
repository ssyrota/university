FROM mambaorg/micromamba:latest
WORKDIR /app
COPY environment.yml .
RUN micromamba create -f environment.yml
COPY . .
RUN eval "$(micromamba shell hook --shell )"
CMD ["/bin/bash", "-c", "micromamba activate lab1_multimedia; sleep 100000"]