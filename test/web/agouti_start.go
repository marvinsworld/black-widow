package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"fmt"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		//var agoutiDriver *agouti.WebDriver
		var err error
		agoutiDriver := agouti.ChromeDriver()
		agoutiDriver.Start()
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		//Expect(page.Destroy()).To(Succeed())
	})

	It("should manage user authentication", func() {
		By("redirecting the user to the login form from the home page", func() {
			Expect(page.Navigate("http://localhost:9515/")).To(Succeed())
			//Expect(page).To(HaveURL("http://localhost:3000/login"))
		})

		By("allowing the user to fill out the login form and submit it", func() {
			loginPrompt, _ := page.FindByID("u1").Text()

			fmt.Println("browser"+loginPrompt)

			//Eventually(page.FindByLabel("E-mail")).Should(BeFound())
			//Expect(page.Find("#account").Fill("payadmin")).To(Succeed())
			//Expect(page.Find("#pwd").Fill("Payadmin123")).To(Succeed())
			//Expect(page.FindByLabel("Password").Fill("secret-password")).To(Succeed())
			//Expect(page.Find("#btnLogin").Click()).To(Succeed())
			//Expect(page.Find("#login_form").Submit()).To(Succeed())
		})
		//
		//By("allowing the user to view their profile", func() {
		//	Eventually(page.FindByLink("Profile Page")).Should(BeFound())
		//	Expect(page.FindByLink("Profile Page").Click()).To(Succeed())
		//	profile := page.Find("section.profile")
		//	Eventually(profile.Find(".greeting")).Should(HaveText("Hello Spud!"))
		//	Expect(profile.Find("img#profile_pic")).To(BeVisible())
		//})
		//
		//By("allowing the user to log out", func() {
		//	Expect(page.Find("#logout").Click()).To(Succeed())
		//	Expect(page).To(HavePopupText("Are you sure?"))
		//	Expect(page.ConfirmPopup()).To(Succeed())
		//	Eventually(page).Should(HaveTitle("Login"))
		//})
	})
})
