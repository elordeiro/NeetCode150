from collections import defaultdict
import heapq

class Account:
    
    def __init__(self, timestamp: int, account_id: str):
        self.creation_time = timestamp
        self.account_id = account_id
        self.transactions = {}
        self.balance = 0
        self.total_outgoing = (0, 0)
        self.schedule = {}
    
    def deposit(self, timestamp: int, amount: int) -> int:
        self.transactions[timestamp] = amount
        self.balance += amount
        return self.balance
    
    def withdraw(self, timestamp: int, amount: int) -> int:
        if self.balance < amount:
            return None
        self.transactions[timestamp] = -amount
        self.balance -= amount
        
        total_outgoing_transactions, outgoing_total = self.total_outgoing
        self.total_outgoing = (total_outgoing_transactions + 1, outgoing_total + amount)
        
        return self.balance


class BankingSystem():

    def __init__(self):
        self.accounts = {}
        self.schedule = defaultdict(list)

    def create_account(self, timestamp: int, account_id: str) -> bool:
        self.check_schedule(timestamp)
        if account_id in self.accounts:
            return False
        self.accounts[account_id] = Account(timestamp, account_id)
        return True
    
    def deposit(self, timestamp: int, account_id: str, amount: int) -> int:
        self.check_schedule(timestamp)
        if account_id not in self.accounts:
            return None
        
        return self.accounts[account_id].deposit(timestamp, amount)
    
    def transfer(self, timestamp: int, source_account_id: str, target_account_id: str, amount: int) -> int:
        self.check_schedule(timestamp)
        if (source_account_id not in self.accounts or 
            target_account_id not in self.accounts or 
            source_account_id == target_account_id):
            return None
        
        if self.accounts[source_account_id].balance < amount:
            return None
        
        self.accounts[target_account_id].deposit(timestamp, amount)
        return self.accounts[source_account_id].withdraw(timestamp, amount)
        
    def top_spenders(self, timestamp: int, n: int) -> list[str]:
        self.check_schedule(timestamp)
        all_accounts = []
        for account in self.accounts.values():
            transactions, amount = account.total_outgoing
            all_accounts.append((-amount, account.account_id))
        
        result = []
        heapq.heapify(all_accounts)
        while n > 0 and all_accounts:
            amount, account_id = heapq.heappop(all_accounts)
            result.append(f"{account_id}({-amount})")
            n -= 1
        
        return result
    
    def schedule_payment(self, timestamp: int, account_id: str, amount: int, delay: int) -> str:
        self.check_schedule(timestamp)
        if account_id not in self.accounts:
            return None
        
        payment_id = "payment" + str(len(self.accounts[account_id].schedule) + 1)
        self.schedule[timestamp + delay].append((account_id, payment_id))
        self.accounts[account_id].schedule[payment_id] = (amount, timestamp + delay)
        result = []
        for payment in self.accounts[account_id].schedule:
            result.append(payment)
        
        return ",".join(result)
    
    def check_schedule(self, timestamp: int):
        if timestamp not in self.schedule:
            return
        
        while self.schedule[timestamp]:
            account_id, payment_id = self.schedule[timestamp].pop(0)
            amount, _ = self.accounts[account_id].schedule[payment_id]
            self.accounts[account_id].withdraw(timestamp, amount)
            self.accounts[account_id].schedule.pop(payment_id)
        
    def cancel_payment(self, timestamp: int, account_id: str, payment_id: str) -> bool:
        self.check_schedule(timestamp)
        if account_id not in self.accounts:
            return False

        if payment_id in self.accounts[account_id].schedule:
            _, timestamp = self.accounts[account_id].schedule.pop(payment_id)
            for idx, (acc_id, p_id) in enumerate(self.schedule[timestamp]):
                if account_id == acc_id and p_id == payment_id:
                    break
            self.schedule[timestamp].pop(idx)
            return True
        return False
            
                
bs = BankingSystem()
print(bs.create_account(1, "account1"))
print(bs.create_account(2, "account2"))
print(bs.deposit(3, "account1", 2000))
print(bs.deposit(4, "account2", 3000))
print(bs.schedule_payment(5, "account1", 50, 10)) # time = 15
print(bs.schedule_payment(6, "account2", 1000, 5)) # time = 11
print(bs.schedule_payment(7, "account1", 3000, 7)) # time = 14
print(bs.deposit(11, "account2", 5))
print(bs.cancel_payment(12, "account2", "payment1"))
print(bs.cancel_payment(13, "account1", "payment1"))
print(bs.deposit(14, "account1", 5))
print(bs.deposit(15, "account1", 5))