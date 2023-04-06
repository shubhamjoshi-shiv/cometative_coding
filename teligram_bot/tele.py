import os
import telebot


API_KEY = "5543511735:AAF1NnRdL1Psx12O3QM3thBLD0dT6pzwXfA"
cid = "1487751384"
bot = telebot.TeleBot(API_KEY)
bot.send_message(cid, "hey man!")


bot.polling()