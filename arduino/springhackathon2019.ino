#define trigPin 12
#define echoPin 13
void setup()
{
  Serial.begin (9600);
  pinMode(trigPin, OUTPUT);
  pinMode(echoPin, INPUT);
}

void loop()
{
  float distance, duration;
  const float total = 47.0;
  int capacity;
  digitalWrite(trigPin, LOW);
  delayMicroseconds(5);
  digitalWrite(trigPin, HIGH);
  delayMicroseconds(10);
  digitalWrite(trigPin, LOW);
  duration = pulseIn(echoPin, HIGH);
  distance = (duration/2) / 29.5;
  // Serial.print(distance);
  // Serial.println(" cm");
  capacity = ((total - distance) / total) * 100;
  if (capacity <= 0) // If truckload is empty
  {
    capacity = 0;
    Serial.print(capacity);
    Serial.println("% Capacity");
    Serial.println("WARNING: EMPTY CAPACITY");
  }
  else if (capacity == 100) // If truckload is full
  {
    Serial.print(capacity);
    Serial.println("% Capacity");
    Serial.println("WARNING: AT CAPACITY");
  }
  else if (capacity > 100) // If truckload has an overflow
  {
    Serial.print(capacity);
    Serial.println("% Capacity");
    Serial.println("WARNING: CAPACITY OVERFLOW");
  }
  else // Everything else about truckload capacity
  {
    Serial.print(capacity);
    Serial.println("% Capacity");
  }
  delay(2000);
} 
