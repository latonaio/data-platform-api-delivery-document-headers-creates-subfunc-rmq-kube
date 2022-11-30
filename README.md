# data-platform-api-invoice-document-headers-creates-subfunc-rmq-kube
data-platform-api-invoice-document-headers-creates-subfunc-rmq-kube は、データ連携基盤において、請求APIサービスのヘッダ登録補助機能を担うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 対象APIサービス
data-platform-api-invoice-document-headers-creates-subfunc-rmq-kube の 対象APIサービスは次の通りです。

*  APIサービス URL: https://xxx.xxx.io/api/API_INVOICE_DOCUMENT_SRV/creates/

## 本レポジトリ が 対応する データ
data-platform-api-invoice-document-headers-creates-subfunc-rmq-kube が対応する データ は、次のものです。

* InvoiceDocument（請求 - ヘッダデータ）
* InvoiceDocumentHeaderPartner（請求 - ヘッダ取引先データ）
* InvoiceDocumentHeaderPartnerContact（請求 - ヘッダ取引先コンタクトデータ）

## Output
data-platform-api-delivery-document-headers-creates-subfunc-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。以下の項目のうち、"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
{
  "connection_key": "request",
  "result": true,
  "redis_key": "abcdefg",
  "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
  "api_status_code": 200,
  "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
  "business_partner": 201,
  "service_label": "INVOICE_DOCUMENT",
  "message": {
    "Header": [
      {
        "InvoiceDocument": 90000004,
        "CreationDate": "",
        "LastChangeDate": "",
        "BillToParty": null,
        "BillFromParty": null,
        "BillToCountry": "JP",
        "BillFromCountry": "US",
        "InvoiceDocumentDate": "",
        "InvoiceDocumentTime": "",
        "InvoicePeriodStartDate": "",
        "InvoicePeriodEndDate": "",
        "AccountingPostingDate": "",
        "InvoiceDocumentIsCancelled": null,
        "CancelledInvoiceDocument": null,
        "IsExportImportDelivery": null,
        "HeaderBillingIsConfirmed": null,
        "HeaderBillingConfStatus": "",
        "TotalNetAmount": 10000,
        "TotalTaxAmount": 1000,
        "TotalGrossAmount": 11000,
        "TransactionCurrency": "JPY",
        "Incoterms": "CIF",
        "PaymentTerms": "0001",
        "DueCalculationBaseDate": null,
        "NetPaymentDays": null,
        "PaymentMethod": "T",
        "HeaderPaymentBlockStatus": null,
        "ExternalReferenceDocument": "",
        "DocumentHeaderText": ""
      }
    ],
    "HeaderPartner": [
      {
        "InvoiceDocument": 90000004,
        "PartnerFunction": "BUYER",
        "BusinessPartner": 101,
        "BusinessPartnerFullName": "株式会社ABC本社",
        "BusinessPartnerName": "ABC本社",
        "Organization": null,
        "Country": "",
        "Language": "JA",
        "Currency": "JPY",
        "ExternalDocumentID": null,
        "AddressID": 100000
      },
      {
        "InvoiceDocument": 90000004,
        "PartnerFunction": "DELIVERTO",
        "BusinessPartner": 102,
        "BusinessPartnerFullName": "株式会社ABC虎ノ門店",
        "BusinessPartnerName": "ABC虎ノ門店",
        "Organization": null,
        "Country": "",
        "Language": "JA",
        "Currency": "JPY",
        "ExternalDocumentID": null,
        "AddressID": 200000
      },
      {
        "InvoiceDocument": 90000004,
        "PartnerFunction": "SELLER",
        "BusinessPartner": 201,
        "BusinessPartnerFullName": "パン販売株式会社",
        "BusinessPartnerName": "パン販売",
        "Organization": null,
        "Country": "",
        "Language": "JA",
        "Currency": "JPY",
        "ExternalDocumentID": null,
        "AddressID": 300000
      }
    ]
  },
  "api_schema": "DPFMInvoiceDocumentCreates",
  "accepter": [
    "All"
  ],
  "deleted": false,
  "sql_update_result": null,
  "sql_update_error": "",
  "subfunc_result": true,
  "subfunc_error": "",
  "exconf_result": null,
  "exconf_error": "",
  "api_processing_result": null,
  "api_processing_error": ""
}
```