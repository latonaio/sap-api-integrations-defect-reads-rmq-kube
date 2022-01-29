package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-defect-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			DefectInternalID:            data.DefectInternalID,
			Defect:                      data.Defect,
			DefectCategory:              data.DefectCategory,
			CreationDate:                data.CreationDate,
			LastChangeDate:              data.LastChangeDate,
			DefectText:                  data.DefectText,
			DefectCodeCatalog:           data.DefectCodeCatalog,
			DefectCodeGroup:             data.DefectCodeGroup,
			DefectCode:                  data.DefectCode,
			DefectCodeVersion:           data.DefectCodeVersion,
			DefectObjectCodeCatalog:     data.DefectObjectCodeCatalog,
			DefectObjectCodeGroup:       data.DefectObjectCodeGroup,
			DefectObjectCode:            data.DefectObjectCode,
			DefectiveQuantity:           data.DefectiveQuantity,
			DefectiveQuantityUnit:       data.DefectiveQuantityUnit,
			ManufacturingOrder:          data.ManufacturingOrder,
			OrderInternalID:             data.OrderInternalID,
			ManufacturingOrderOperation: data.ManufacturingOrderOperation,
			ManufacturingOrderSequence:  data.ManufacturingOrderSequence,
			CreationTime:                data.CreationTime,
			LastChangeTime:              data.LastChangeTime,
			DefectClass:                 data.DefectClass,
			NumberOfDefects:             data.NumberOfDefects,
			InspPlanOperationInternalID: data.InspPlanOperationInternalID,
			InspectionCharacteristic:    data.InspectionCharacteristic,
			InspectionSubsetInternalID:  data.InspectionSubsetInternalID,
			MaterialSample:              data.MaterialSample,
			WorkCenterTypeCode:          data.WorkCenterTypeCode,
			MainWorkCenterInternalID:    data.MainWorkCenterInternalID,
			MainWorkCenterPlant:         data.MainWorkCenterPlant,
			MainWorkCenter:              data.MainWorkCenter,
			Equipment:                   data.Equipment,
			FunctionalLocation:          data.FunctionalLocation,
			IsDeleted:                   data.IsDeleted,
			DefectOrigin:                data.DefectOrigin,
			Material:                    data.Material,
			Plant:                       data.Plant,
			InspectionLot:               data.InspectionLot,
			CatalogProfile:              data.CatalogProfile,
			ChangedDateTime:             data.ChangedDateTime,
		})
	}

	return header, nil
}
