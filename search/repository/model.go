package repository

type ResponseStruct struct {
    Status string `json:"status"`
    Result ResultTitle `json:"result"`
}

type ResultTitle struct {
    Title string `json:"title"`
}