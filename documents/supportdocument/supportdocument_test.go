package supportdocument

import (
	"strings"
	"testing"
	"time"

	"github.com/diegofxm/ubl21-dian/documents/common/model"
	"github.com/diegofxm/ubl21-dian/documents/common/types"
)

func TestSupportDocumentBuilder(t *testing.T) {
	t.Run("Build SupportDocument", func(t *testing.T) {
		// Datos de prueba
		buyer := model.PartyData{
			PersonType:   "1",
			ID:           "900123456",
			DV:           "1",
			DocumentType: "31",
			Name:         "MI EMPRESA SAS",
			TaxLevelCode: "O-13",
			TaxSchemeID:  "01",
			TaxSchemeName: "IVA",
			Address: model.AddressData{
				ID:                   "11001",
				CityName:             "Bogotá",
				CountrySubentity:     "Bogotá",
				CountrySubentityCode: "11",
				AddressLine:          "Calle 123 # 45-67",
				Country: model.CountryData{
					Code: "CO",
					Name: "Colombia",
				},
			},
		}
		
		supplier := model.PartyData{
			PersonType:   "1",
			ID:           "800654321",
			DV:           "9",
			DocumentType: "31",
			Name:         "PROVEEDOR XYZ SAS",
			TaxLevelCode: "O-13",
			TaxSchemeID:  "01",
			TaxSchemeName: "IVA",
			Address: model.AddressData{
				ID:                   "11001",
				CityName:             "Bogotá",
				CountrySubentity:     "Bogotá",
				CountrySubentityCode: "11",
				AddressLine:          "Carrera 10 # 20-30",
				Country: model.CountryData{
					Code: "CO",
					Name: "Colombia",
				},
			},
		}
		
		// Línea de documento
		line := InvoiceLineXML{
			ID: types.CBCElement{Value: "1"},
			InvoicedQuantity: types.QuantityElement{
				UnitCode: "EA",
				Value:    "10",
			},
			LineExtensionAmount: types.AmountElement{
				CurrencyID: "COP",
				Value:      "100000.00",
			},
			Item: ItemXML{
				Description: []types.CBCElement{
					{Value: "Producto de prueba"},
				},
			},
			Price: PriceXML{
				PriceAmount: types.AmountElement{
					CurrencyID: "COP",
					Value:      "10000.00",
				},
			},
		}
		
		// Totales
		totals := LegalMonetaryTotalXML{
			LineExtensionAmount: types.AmountElement{
				CurrencyID: "COP",
				Value:      "100000.00",
			},
			TaxExclusiveAmount: types.AmountElement{
				CurrencyID: "COP",
				Value:      "100000.00",
			},
			TaxInclusiveAmount: types.AmountElement{
				CurrencyID: "COP",
				Value:      "119000.00",
			},
			PayableAmount: types.AmountElement{
				CurrencyID: "COP",
				Value:      "119000.00",
			},
		}
		
		// Construir documento soporte
		builder := NewBuilder()
		xml, err := builder.
			SetID("DS001").
			SetCUDS("abc123def456cuds").
			SetIssueDate("2025-01-31").
			SetIssueTime("14:30:00-05:00").
			SetProfileExecutionID("1").
			AddNote("Documento soporte de compra").
			SetBuyer(buyer).
			SetSupplier(supplier).
			AddLine(line).
			SetLegalMonetaryTotal(totals).
			Build()
		
		if err != nil {
			t.Fatalf("Error building SupportDocument: %v", err)
		}
		
		// Verificaciones
		if !strings.Contains(xml, "<Invoice") {
			t.Error("XML should contain Invoice root element")
		}
		
		if !strings.Contains(xml, "DS001") {
			t.Error("XML should contain document ID")
		}
		
		if !strings.Contains(xml, "abc123def456cuds") {
			t.Error("XML should contain CUDS")
		}
		
		if !strings.Contains(xml, "<cbc:InvoiceTypeCode>05</cbc:InvoiceTypeCode>") {
			t.Error("XML should contain InvoiceTypeCode 05")
		}
		
		if !strings.Contains(xml, "MI EMPRESA SAS") {
			t.Error("XML should contain buyer name")
		}
		
		if !strings.Contains(xml, "PROVEEDOR XYZ SAS") {
			t.Error("XML should contain supplier name")
		}
		
		if !strings.Contains(xml, "Producto de prueba") {
			t.Error("XML should contain item description")
		}
		
		if !strings.Contains(xml, "119000.00") {
			t.Error("XML should contain payable amount")
		}
		
		t.Log("✓ SupportDocument XML generated successfully")
	})
	
	t.Run("Build with billing reference", func(t *testing.T) {
		builder := NewBuilder()
		
		// Agregar referencia a factura de proveedor
		builder.AddBillingReference(model.BillingReferenceData{
			InvoiceID: "FV-12345",
			IssueDate: time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
		})
		
		model := builder.GetModel()
		
		if len(model.BillingReference) != 1 {
			t.Errorf("Expected 1 billing reference, got %d", len(model.BillingReference))
		}
		
		if model.BillingReference[0].InvoiceDocumentReference.ID.Value != "FV-12345" {
			t.Errorf("Expected invoice ID FV-12345, got %s", 
				model.BillingReference[0].InvoiceDocumentReference.ID.Value)
		}
		
		t.Log("✓ Billing reference added correctly")
	})
	
	t.Run("Build with withholding taxes", func(t *testing.T) {
		builder := NewBuilder()
		
		// Agregar retención
		withholding := types.TaxTotalXML{
			TaxAmount: types.AmountElement{
				CurrencyID: "COP",
				Value:      "2500.00",
			},
			TaxSubtotal: []types.TaxSubtotalXML{
				{
					TaxableAmount: types.AmountElement{
						CurrencyID: "COP",
						Value:      "100000.00",
					},
					TaxAmount: types.AmountElement{
						CurrencyID: "COP",
						Value:      "2500.00",
					},
					Percent: types.CBCElement{Value: "2.5"},
					TaxCategory: types.TaxCategoryXML{
						TaxScheme: types.TaxSchemeXML{
							ID:   types.CBCElement{Value: "07"},
							Name: types.CBCElement{Value: "Retención en la fuente"},
						},
					},
				},
			},
		}
		
		builder.AddWithholdingTaxTotal(withholding)
		
		model := builder.GetModel()
		
		if len(model.WithholdingTaxTotal) != 1 {
			t.Errorf("Expected 1 withholding tax, got %d", len(model.WithholdingTaxTotal))
		}
		
		if model.WithholdingTaxTotal[0].TaxAmount.Value != "2500.00" {
			t.Errorf("Expected withholding amount 2500.00, got %s", 
				model.WithholdingTaxTotal[0].TaxAmount.Value)
		}
		
		t.Log("✓ Withholding tax added correctly")
	})
}
