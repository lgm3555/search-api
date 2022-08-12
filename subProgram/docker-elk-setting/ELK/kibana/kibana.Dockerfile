ARG ELK_VERSION

# https://www.docker.elastic.co/
FROM docker.elastic.co/kibana/kibana:${ELK_VERSION}

# 여기에 PLUGIN 추가 코드
# Example: RUN kibana-plugin install <name|url>