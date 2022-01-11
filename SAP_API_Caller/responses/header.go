package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			DefectInternalID            string `json:"DefectInternalID"`
			Defect                      string `json:"Defect"`
			DefectCategory              string `json:"DefectCategory"`
			CreationDate                string `json:"CreationDate"`
			LastChangeDate              string `json:"LastChangeDate"`
			DefectText                  string `json:"DefectText"`
			DefectCodeCatalog           string `json:"DefectCodeCatalog"`
			DefectCodeGroup             string `json:"DefectCodeGroup"`
			DefectCode                  string `json:"DefectCode"`
			DefectCodeVersion           string `json:"DefectCodeVersion"`
			DefectObjectCodeCatalog     string `json:"DefectObjectCodeCatalog"`
			DefectObjectCodeGroup       string `json:"DefectObjectCodeGroup"`
			DefectObjectCode            string `json:"DefectObjectCode"`
			DefectiveQuantity           string `json:"DefectiveQuantity"`
			DefectiveQuantityUnit       string `json:"DefectiveQuantityUnit"`
			ManufacturingOrder          string `json:"ManufacturingOrder"`
			OrderInternalID             string `json:"OrderInternalID"`
			ManufacturingOrderOperation string `json:"ManufacturingOrderOperation"`
			ManufacturingOrderSequence  string `json:"ManufacturingOrderSequence"`
			CreationTime                string `json:"CreationTime"`
			LastChangeTime              string `json:"LastChangeTime"`
			DefectClass                 string `json:"DefectClass"`
			NumberOfDefects             int    `json:"NumberOfDefects"`
			InspPlanOperationInternalID string `json:"InspPlanOperationInternalID"`
			InspectionCharacteristic    string `json:"InspectionCharacteristic"`
			InspectionSubsetInternalID  string `json:"InspectionSubsetInternalID"`
			MaterialSample              string `json:"MaterialSample"`
			WorkCenterTypeCode          string `json:"WorkCenterTypeCode"`
			MainWorkCenterInternalID    string `json:"MainWorkCenterInternalID"`
			MainWorkCenterPlant         string `json:"MainWorkCenterPlant"`
			MainWorkCenter              string `json:"MainWorkCenter"`
			Equipment                   string `json:"Equipment"`
			FunctionalLocation          string `json:"FunctionalLocation"`
			IsDeleted                   string `json:"IsDeleted"`
			DefectOrigin                string `json:"DefectOrigin"`
			Material                    string `json:"Material"`
			Plant                       string `json:"Plant"`
			InspectionLot               string `json:"InspectionLot"`
			CatalogProfile              string `json:"CatalogProfile"`
			ChangedDateTime             string `json:"ChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
