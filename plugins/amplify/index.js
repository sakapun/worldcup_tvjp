import amplify from "aws-amplify";

import AppSyncConfig from "./AppSync";

let myAppConfig = {
  'aws_appsync_graphqlEndpoint': AppSyncConfig.graphqlEndpoint,
  'aws_appsync_region': AppSyncConfig.region,
  'aws_appsync_authenticationType': AppSyncConfig.authenticationType,
  'aws_appsync_apiKey': AppSyncConfig.apiKey,
};

amplify.configure(myAppConfig);

export default amplify;
