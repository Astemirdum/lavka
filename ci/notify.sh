#!/bin/bash

echo "${1}"

TG_URL="https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage"
MSG_TEXT="CI/CD status: $1\n Project: $CI_PROJECT_NAME\nURL: $CI_PROJECT_URL/pipelines/$CI_PIPELINE_ID/\nBranch: $CI_COMMIT_REF_SLUG"

curl -s --max-time 5 -X POST -H "Content-Type: application/json" \
  -d "{\"chat_id\":\"${TELEGRAM_USER_ID}\",\"disable_web_page_preview\":1,\"text\":\"${MSG_TEXT}\"}" "${TG_URL}" > /dev/null

