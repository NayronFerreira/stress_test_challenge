package constants

import "time"

const DefaultRequestTimeout = 10 * time.Second

const (
	ReportHeader = "=========================================\n                 RELATORY\n========================================="
	ReportFooter = "========================================="
)

const StressTestAsciiArt = `
    SSSS   TTTTT  RRRR   EEEE  SSSS  SSSS   TTTTT  EEEE  SSSS  TTTTT
   S        T    R   R  E    S      S        T    E    S        T  
    SSS     T    RRRR   EEE   SSS    SSS     T    EEE   SSS     T  
       S    T    R  R   E        S      S    T    E        S    T  
   SSSS     T    R   R  EEEE  SSSS   SSSS    T    EEEE  SSSS    T  
    `
