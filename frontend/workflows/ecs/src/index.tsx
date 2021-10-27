import type { BaseWorkflowProps, WorkflowConfiguration } from "@clutch-sh/core";

import HelloWorld from "./hello-world";

export interface WorkflowProps extends BaseWorkflowProps {}

const register = (): WorkflowConfiguration => {
  return {
    developer: {
      name: "Joe",
      contactUrl: "mailto:josephveejay@gmail.com",
    },
    path: "ecs",
    group: "ECS",
    displayName: "ECS",
    routes: {
      landing: {
        path: "/",
        description: "Lists ECS Clusters.",
        component: HelloWorld,
      },
    },
  };
};

export default register;
