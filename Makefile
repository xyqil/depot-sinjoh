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

%.so:%.so.o
  $(CXX) -shared $@ -o $<
%.a:%.a.o
  $(AR) rvs $< $@
%.cpp:%.py2.py
  $(PYTHON) -m cython $@ --embed 
main$(EXE_EXTENSION):$(LIBS) $(SOS) $(OBJS)
  $(CXX) -o $@ $^ $(CXXFLAGS)
