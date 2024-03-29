import React from 'react';
import TabbedPageContainer from 'components/TabbedPageContainer';
import DetailsPageWrapper from 'components/DetailsPageWrapper';
import { RESOURCE_TYPES } from '../utils';
import TabDetails from './TabDetails';
import TabCisBenchmark from './TabCisBenchmark';

import './application-resources-details.scss';

const DetailsContent = ({data}) => {
    const {applicationResource} = data || {};
    const {id, resourceType} = applicationResource || {};

    const isImage = resourceType === RESOURCE_TYPES.IMAGE.value;

    return (
        <TabbedPageContainer
            items={[
                {
                    id: "details",
                    title: "Details",
                    isIndex: true,
                    component: () => <TabDetails data={data} />
                },
                {
                    id: "cis",
                    title: "CIS Docker Benchmark",
                    path: "cisbenchmark",
                    component: () => <TabCisBenchmark id={id} />,
                    disabled: !isImage,
                    tabTooltip: isImage ? null : "CIS Docker benchmark is only available for images"
                }
            ]}
        />
    )
}

const ApplicationResourcesDetails = () => (
    <DetailsPageWrapper
        title="Application resource information"
        backTitle="Application resources"
        url="applicationResources"
        detailsContent={DetailsContent}
        getReplace={params => {
            const id = params["id"];
            const innerTab = params["*"];
            
            return !!innerTab ? `/${id}/${innerTab}` : `/${id}`;
        }}
    />
)

export default ApplicationResourcesDetails;