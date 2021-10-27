import React from "react";
import AccountTreeIcon from "@material-ui/icons/AccountTree";
import { Wizard, WizardStep } from "@clutch-sh/wizard";
import type { WizardChild } from "@clutch-sh/wizard";
import { Button, 
  ButtonGroup, 
  client, 
  useWizardContext, 
  Table, 
  TableRow,
  Select, TableRowActions, TableRowAction } from "@clutch-sh/core";
import { useDataLayout } from "@clutch-sh/data-layout";
import type { WorkflowProps } from ".";

const ECSLookup: React.FC<WorkflowProps> = ({ heading }) => {
  const { onSubmit } = useWizardContext();
  return (
    <>
    <Select
      label="Select Environment"
      name="storybookDemo"
      onChange={function noRefCheck(){}}
      options={[
        {label: 'STAG'},
        {label: 'PROD'}
      ]}
    />
    <ButtonGroup>
      <Button text="Clusters" onClick={onSubmit}/>
    </ButtonGroup>
    </>
  );
};


const ECSClusters: React.FC<WorkflowProps> = ({ heading }) => {
  const dataLayout = {
    clusterData: {
      deps: [],
      hydrator: () => {
        return client
          .post("/v1/aws/ecs/getClusters", {
            region: "eu-west-1",
          })
          .then(response => {
            console.log(response)
            return response?.data?.cluster || [];
          });
      },
    },
  };

  return (
    <>
    <Wizard dataLayout={dataLayout} heading={heading}>
    <ECSLookup name="Clusters"/>
    <ECSDetails name="Services"/>
    </Wizard>
    </>
  );
};

const ECSDetails: React.FC<WizardChild> = () => {
  const clusterData = useDataLayout("clusterData");
  let ecsResults = clusterData.displayValue();
  if (_.isEmpty(ecsResults)) {
    ecsResults = [];
  }

  return (
    <WizardStep error={clusterData.error} isLoading={clusterData.isLoading}>
      <Table actionsColumn columns={["Clusters"]} >
        {ecsResults.map((cluster, index: number) => (
        <TableRow key={index}>
          {cluster}
          <TableRowActions>
          <TableRowAction icon={<AccountTreeIcon />} onClick={function noRefCheck(){}}>
          Services
          </TableRowAction>
          </TableRowActions>
        </TableRow>
          ))}
      </Table>
    </WizardStep>
  );
};


export default ECSClusters;
