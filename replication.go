import (
    "fmt"
    "github.com/bndr/gojenkins"
)

oldJenkins, err := gojenkins.CreateJenkins(nil, "http://old-jenkins-url", "username", "password").Init()
if err != nil {
    fmt.Println("Failed to connect to old Jenkins:", err)
    return
}

newJenkins, err := gojenkins.CreateJenkins(nil, "http://new-jenkins-url", "username", "password").Init()
if err != nil {
    fmt.Println("Failed to connect to new Jenkins:", err)
    return
}




oldJobs, err := oldJenkins.GetAllJobs()
if err != nil {
    fmt.Println("Failed to retrieve jobs from old Jenkins:", err)
    return
}



for _, oldJob := range oldJobs {
    newJob := newJenkins.InitJob(oldJob.Name)

    // Customize job parameters if needed
    // newJob.SetParameter("param_name", "param_value")

    _, err := newJob.Build()
    if err != nil {
        fmt.Printf("Failed to create job %s on new Jenkins: %v\n", oldJob.Name, err)
    } else {
        fmt.Printf("Successfully created job %s on new Jenkins\n", oldJob.Name)
    }
}
