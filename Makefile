UNAME := $(shell uname)
DO_ERROR := yes
EXE_EXTENSION =
ifeq ($(UNAME), Linux)
  SO := so
  DO_ERROR := no
endif
#ifneq (,findstring ($(UNAME), Revolution))
#  $(error How the heck did you manage to get GNU/Make *and* uname ported to Wii?)
#endif

ifeq ($(DO_ERROR), yes)
  $(error Unsupported platform $(UNAME))
endif
LIBS := 
SOS :=
SOURCES := 
OBJS := $(addsuffix .o, $(basename $(notdir $(SOURCES))))
CXXFLAGS :=
CXX := g++
GO := go
PYTHON := python3

%.o:%.cpp
	$(CXX) $(CXXFLAGS) -c -o $@ $<
main$(EXE_EXTENSION):$(LIBS) $(SOS) $(OBJS)
  $(CXX) -o $@ $^ $(CXXFLAGS)
