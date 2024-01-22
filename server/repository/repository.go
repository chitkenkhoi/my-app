package repository

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/model"
	Mongo "server/mongo"
)

//	func CheckId(id string) bool {
//		client := FB.Client
//		doc, err := client.Collection("users").Doc(id).Get(context.Background())
//		if err != nil {
//			log.Fatalf("Error getting document: %v", err)
//			return false
//		}
//		if doc.Exists() {
//			return true
//		} else {
//			return false
//		}
//	}
//
//	func GetItem(id string) *model.Item {
//		client := FB.Client
//		doc, err := client.Collection("items").Doc(id).Get(context.Background())
//		if err != nil {
//			return nil
//		}
//		var item model.Item
//		doc.DataTo(&item)
//		item.ID = id
//		return &item
//	}
//
//	func CheckUser(user model.User) string {
//		client := FB.Client
//		query := client.Collection("users").
//			Where("username", "==", user.Username).
//			Where("password", "==", user.Password).
//			Documents(context.Background())
//
//		// Check if there is any matching document
//		for {
//			doc, err := query.Next()
//			if err == iterator.Done {
//				break
//			}
//			if err != nil {
//				log.Fatalf("Error querying documents: %v", err)
//			}
//			if doc != nil {
//				return doc.Ref.ID
//			}
//			// Document with the specified username and password exists
//
//		}
//		return ""
//	}
//
//	func GetItems(name string) []model.Item {
//		client := FB.Client
//		query := client.Collection("items").
//			Where("name", "==", name).
//			Documents(context.Background())
//		// Iterate through the result set
//		var items []model.Item
//		for {
//			doc, err := query.Next()
//			if err == iterator.Done {
//				break
//			}
//			if err != nil {
//				log.Fatalf("Error querying documents: %v", err)
//			}
//			var item model.Item
//			doc.DataTo(&item)
//			item.ID = doc.Ref.ID
//			items = append(items, item)
//		}
//		return items
//	}
//
//	func InsertItem(item model.Item) error {
//		client := FB.Client
//		_, _, err := client.Collection("items").Add(context.Background(), item)
//		if err != nil {
//			log.Fatalf("Error adding document: %v", err)
//			return err
//		}
//		fmt.Println("Document added successfully!")
//		return nil
//	}
//
//	func UpdateItem(item model.Item) error {
//		client := FB.Client
//		id := item.ID
//		doc, err := client.Collection("items").Doc(id).Get(context.Background())
//		if err != nil {
//			return err
//		}
//		if _, er := doc.Ref.Update(context.Background(), []firestore.Update{
//			{Path: "name", Value: item.Name},
//			{Path: "price", Value: item.Price},
//			{Path: "size", Value: item.Size},
//			{Path: "unit", Value: item.Unit},
//		}); er != nil {
//			return er
//		}
//
//		fmt.Println("Item updated successfully!")
//		return nil
//	}
//
//	func DeleteItem(id string) error {
//		client := FB.Client
//		doc, err := client.Collection("items").Doc(id).Get(context.Background())
//		if err != nil {
//			return err
//		}
//		doc.Ref.Delete(context.Background())
//		return nil
//	}
//
//	func CountTotalOrders(filter int) int {
//		client := FB.Client
//		count := 0
//		collection := client.Collection("orders")
//		var query *firestore.DocumentIterator
//		if filter == 1 {
//			query = collection.Where("status1", "==", false).Documents(context.Background())
//		} else if filter == 2 {
//			query = collection.Where("status2", "==", false).Documents(context.Background())
//		} else if filter == 3 {
//			query = collection.Where("status1", "==", true).Where("status2", "==", true).Documents(context.Background())
//		} else {
//			query = collection.Documents(context.Background())
//		}
//
//		for {
//			_, err := query.Next()
//			if err == iterator.Done {
//				break
//			}
//			if err != nil {
//				log.Fatalf("Error count total orders: %v", err)
//			}
//
//			count++
//		}
//		return count
//	}
//
//	func CountSomeOrders(data string, filter int) int {
//		client := FB.Client
//		count := 0
//		collection := client.Collection("orders")
//		var query1 *firestore.DocumentIterator
//		var query2 *firestore.DocumentIterator
//		if filter == 1 {
//			query1 = collection.Where("status1", "==", false).Where("buyer", ">=", data).Where("buyer", "<", data+"\uf8ff").Documents(context.Background())
//			query2 = collection.Where("status1", "==", false).Where("buyer", "<=", data).Where("buyer", ">", "\uf8ff"+data).Documents(context.Background())
//		} else if filter == 2 {
//			query1 = collection.Where("status2", "==", false).Where("buyer", ">=", data).Where("buyer", "<", data+"\uf8ff").Documents(context.Background())
//			query2 = collection.Where("status2", "==", false).Where("buyer", "<=", data).Where("buyer", ">", "\uf8ff"+data).Documents(context.Background())
//		} else {
//			query1 = collection.Where("buyer", ">=", data).Where("buyer", "<", data+"\uf8ff").Where("status1", "==", true).Where("status2", "==", true).Documents(context.Background())
//			query2 = collection.Where("buyer", "<=", data).Where("buyer", ">", "\uf8ff"+data).Where("status1", "==", true).Where("status2", "==", true).Documents(context.Background())
//		}
//		for {
//			_, err := query1.Next()
//			if err == iterator.Done {
//				break
//			}
//			if err != nil {
//				log.Fatalf("Error count total orders1: %v", err)
//			}
//			count++
//		}
//		for {
//			_, err := query2.Next()
//			if err == iterator.Done {
//				break
//			}
//			if err != nil {
//				log.Fatalf("Error count total orders: %v", err)
//			}
//			count++
//		}
//		return count
//	}
//
//	func GetOrder(id string) (*model.Order, error) {
//		client := FB.Client
//		docRef := client.Collection("orders").Doc(id)
//		doc, err := docRef.Get(context.Background())
//		if err != nil {
//			return nil, err
//		}
//		var order model.Order
//		doc.DataTo(&order)
//		order.ID = id
//		return &order, nil
//	}
//
//	func InsertOrder(order model.Order) error {
//		client := FB.Client
//		_, _, err := client.Collection("orders").Add(context.Background(), order)
//		if err != nil {
//			log.Fatalf("Error adding order: %v", err)
//			return err
//		}
//		fmt.Println("Order added successfully!")
//		return nil
//	}
//
//	func UpdateOrderStatus(id string, status int) error {
//		client := FB.Client
//		docRef := client.Collection("orders").Doc(id)
//		doc, err := docRef.Get(context.Background())
//		if err != nil {
//			fmt.Println("no id")
//			return err
//		}
//		status1 := doc.Data()["status1"].(bool)
//		status2 := doc.Data()["status2"].(bool)
//		if status == 1 {
//			status1 = !status1
//		} else if status == 2 {
//			status2 = !status2
//		} else {
//			status1 = !status1
//			status2 = !status2
//		}
//		_, err = docRef.Update(context.Background(), []firestore.Update{
//			{Path: "status1", Value: status1},
//			{Path: "status2", Value: status2},
//		})
//		if err != nil {
//			fmt.Println("cant update, sth wrong")
//			return err
//		}
//		fmt.Println("Document updated successfully!")
//		return nil
//	}
//
//	func DeleteOrder(id string) error {
//		client := FB.Client
//		doc, err := client.Collection("orders").Doc(id).Get(context.Background())
//		//_, err := docRef.Delete(context.Background())
//		if err != nil {
//			return err
//		}
//		_, er := doc.Ref.Delete(context.Background())
//		return er
//	}
func Test(ctx *gin.Context) {
	var order model.Order

	database := Mongo.Client.Database("store")
	if err := database.Collection("orders").FindOne(context.Background(), model.Order{Buyer: "The B"}).Decode(&order); err != nil {
		fmt.Println(err)
	}
	fmt.Println(order)
	ctx.JSON(200, gin.H{
		"message": "test",
	})
}
