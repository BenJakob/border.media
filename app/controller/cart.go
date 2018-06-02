package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"borgdir.media/app/model"
)

func Cart(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("cart").ParseFiles("template/cart.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	headerData := model.Header{Title: "Warenkorb", Css: []string{"/css/equipment.css", "/css/style.css"}}
	items := model.GetCartItems(user)
	cartItemCount := model.GetCartItemCount(user)
	footerData := []string{"/scripts/cart.js"}

	cartData := model.CartData{
		HeaderData:    headerData,
		User:          user,
		IsLoggedIn:    true,
		Items:         items,
		CartItemCount: cartItemCount,
		FooterData:    footerData,
	}

	if err := t.Execute(w, cartData); err != nil {
		log.Fatalln(err)
	}
}

func CartUpdateDate(w http.ResponseWriter, r *http.Request) {
	_, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	date := r.FormValue("date")

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	model.UpdateCartItemDate(id, date)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func CartUpdateQuantity(w http.ResponseWriter, r *http.Request) {
	_, err := model.GetLoggedInUser(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	model.UpdateCartItemQuantity(id, quantity)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func CartRemove(w http.ResponseWriter, r *http.Request) {
	_, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	model.DeleteCartItem(id)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func CartCheckout(w http.ResponseWriter, r *http.Request) {
	_, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	user, err := model.GetUser(uint32(id))

	if err != nil {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	model.CheckoutCart(user)
	http.Redirect(w, r, "/my-equipment", http.StatusSeeOther)
}
